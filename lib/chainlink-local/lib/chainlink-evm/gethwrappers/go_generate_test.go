// package gethwrappers_test verifies correct and up-to-date generation of golang wrappers
// for solidity contracts. See go_generate.go for the actual generation.
package gethwrappers

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fatih/color"

	cutils "github.com/smartcontractkit/chainlink-common/pkg/utils"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const compileCommand = "../contracts/scripts/native_solc_compile_all"

// TestCheckContractHashesFromLastGoGenerate compares the abi and bytecode of the
// contract artifacts in contracts/solc with the abi and bytecode stored in the
// contract wrapper
func TestCheckContractHashesFromLastGoGenerate(t *testing.T) {
	tests.SkipShort(t, "requires compiled artifacts")
	versions, err := ReadVersionsDB()
	require.NoError(t, err)
	require.NotEmpty(t, versions.GethVersion, `version DB should have a "GETH_VERSION:" line`)

	wd, err := os.Getwd()
	if err != nil {
		wd = "<directory containing this test>"
	}
	require.Equal(t, versions.GethVersion, GethVersion,
		color.HiRedString(BoxOutput("please re-run `go generate %s` and commit the"+
			"changes", wd)))

	for _, contractVersionInfo := range versions.ContractVersions {
		if isOCRContract(contractVersionInfo.AbiPath) || isVRFV2Contract(contractVersionInfo.AbiPath) {
			continue
		}
		compareCurrentCompilerArtifactAgainstRecordsAndSoliditySources(t, contractVersionInfo)
	}
	// Just check that LinkToken details haven't changed (they never ought to)
	linkDetails, err := os.ReadFile(filepath.Join(getProjectRoot(t), "contracts/LinkToken.json"))
	require.NoError(t, err, "could not read link contract details")
	require.Equal(t, "27c0e17a79553fccc63a4400c6bbe415ff710d9cc7c25757bff0f7580205c922", fmt.Sprintf("%x", sha256.Sum256(linkDetails)),
		"should never differ!")
}

func isOCRContract(fullpath string) bool {
	return strings.Contains(fullpath, "OffchainAggregator")
}

// VRFv2 currently uses revert error types which are not supported by abigen
// and so we have to manually modify the abi to remove them.
func isVRFV2Contract(fullpath string) bool {
	return strings.Contains(fullpath, "VRFCoordinatorV2")
}

// compareCurrentCompilerArtifactAgainstRecordsAndSoliditySources checks that
// the file at each ContractVersion.AbiPath and ContractVersion.BinaryPath hashes to its
// ContractVersion.Hash, and that the solidity source code recorded in the
// compiler artifact matches the current solidity contracts.
//
// Most of the compiler artifacts should contain output from sol-compiler, or
// "pnpm compile". The relevant parts of its schema are
//
//	{ "sourceCodes": { "<filePath>": "<code>", ... } }
//
// where <filePath> is the path to the contract, below the truffle contracts/
// directory, and <code> is the source code of the contract at the time the JSON
// file was generated.
func compareCurrentCompilerArtifactAgainstRecordsAndSoliditySources(
	t *testing.T, versionInfo ContractVersion,
) {
	hash := VersionHash(versionInfo.AbiPath, versionInfo.BinaryPath)
	rootDir := getProjectRoot(t)
	recompileCommand := fmt.Sprintf("(cd %s/contracts; make wrappers-all)", rootDir)
	assert.Equal(t, versionInfo.Hash, hash,
		BoxOutput(`compiled %s and/or %s has changed; please rerun
%s,
and commit the changes`, versionInfo.AbiPath, versionInfo.BinaryPath, recompileCommand))
}

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Short() {
		ensureArtifacts()
	}
	os.Exit(m.Run())
}

// Ensure that solidity compiler artifacts are present before running this test,
// by compiling them if necessary.
func ensureArtifacts() {
	db, err := versionsDBLineReader()
	if err != nil {
		panic(err)
	}
	var solidityArtifactsMissing []string
	for db.Scan() {
		line := strings.Fields(db.Text())
		if stripTrailingColon(line[0], "") != "GETH_VERSION" {
			if os.IsNotExist(cutils.JustError(os.Stat(line[1]))) {
				solidityArtifactsMissing = append(solidityArtifactsMissing, line[1])
			}
		}
	}
	if len(solidityArtifactsMissing) == 0 {
		return
	}
	fmt.Printf("some solidity artifacts missing (%s); rebuilding...",
		solidityArtifactsMissing)
	// Don't want to run "make wrappers-all" here, because that would
	// result in an infinite loop
	cmd := exec.Command("bash", "-c", compileCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// getProjectRoot returns the root of the chainlink project
func getProjectRoot(t *testing.T) (rootPath string) {
	root, err := os.Getwd()
	require.NoError(t, err, "could not get current working directory")
	for root != "/" { // Walk up path to find dir containing go.mod
		if _, err := os.Stat(filepath.Join(root, "go.mod")); !os.IsNotExist(err) {
			return root
		}
		root = filepath.Dir(root)
	}
	t.Fatal("could not find project root")
	return
}
