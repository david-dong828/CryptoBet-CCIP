import React, { useState, useEffect } from 'react';
import { Clock, AlertCircle } from 'lucide-react';
import './CountdownTimer.css';

function CountdownTimer({ endTime, onTimeUp, status }) {
  const [timeLeft, setTimeLeft] = useState(null);
  const [isExpired, setIsExpired] = useState(false);

  useEffect(() => {
    const calculateTimeLeft = () => {
      const now = Date.now();
      const difference = endTime - now;

      if (difference <= 0) {
        setIsExpired(true);
        setTimeLeft({
          days: 0,
          hours: 0,
          minutes: 0,
          seconds: 0
        });
        if (onTimeUp && status === 'Open') {
          onTimeUp();
        }
        return;
      }

      const days = Math.floor(difference / (1000 * 60 * 60 * 24));
      const hours = Math.floor((difference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
      const minutes = Math.floor((difference % (1000 * 60 * 60)) / (1000 * 60));
      const seconds = Math.floor((difference % (1000 * 60)) / 1000);

      setTimeLeft({ days, hours, minutes, seconds });
      setIsExpired(false);
    };

    // Calculate immediately
    calculateTimeLeft();

    // Update every second
    const timer = setInterval(calculateTimeLeft, 1000);

    return () => clearInterval(timer);
  }, [endTime, onTimeUp, status]);

  if (!timeLeft) return null;

  const getUrgencyClass = () => {
    if (isExpired) return 'expired';
    if (timeLeft.days === 0 && timeLeft.hours < 1) return 'urgent';
    if (timeLeft.days === 0 && timeLeft.hours < 24) return 'warning';
    return 'normal';
  };

  const formatTimeUnit = (value, unit) => {
    return (
      <div className="time-unit">
        <div className="time-value">{value.toString().padStart(2, '0')}</div>
        <div className="time-label">{unit}</div>
      </div>
    );
  };

  return (
    <div className={`countdown-timer ${getUrgencyClass()}`}>
      <div className="countdown-header">
        {isExpired ? (
          <>
            <AlertCircle size={20} />
            <span>Market Closed</span>
          </>
        ) : (
          <>
            <Clock size={20} />
            <span>Time Remaining</span>
          </>
        )}
      </div>
      
      {!isExpired && (
        <div className="countdown-display">
          {timeLeft.days > 0 && formatTimeUnit(timeLeft.days, 'Days')}
          {formatTimeUnit(timeLeft.hours, 'Hours')}
          {formatTimeUnit(timeLeft.minutes, 'Min')}
          {formatTimeUnit(timeLeft.seconds, 'Sec')}
        </div>
      )}
      
      {isExpired && status === 'Open' && (
        <div className="countdown-message">
          Waiting for settlement...
        </div>
      )}
    </div>
  );
}

export default CountdownTimer;