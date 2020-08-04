use std::fmt;
extern crate num_integer;
use num_integer::Integer;

#[derive(Debug, PartialEq)]
pub struct Clock {
    hours: i64,
    minutes: i64, 
}

impl Clock {
    pub fn new(hours: i64, minutes: i64) -> Self {
        Clock {
            //Compared to regular '/' and '%',
            //div_floor and mod_floor handle negative numbers 
            hours: (hours * 60 + minutes).div_floor(&60).mod_floor(&24),
            minutes: minutes.mod_floor(&60),
        }
    }

    pub fn add_minutes(self, minutes: i64) -> Self {
        Clock::new(self.hours, self.minutes + minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}
