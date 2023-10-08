#ifndef ZADANIE3_WATCHDOG_H
#define ZADANIE3_WATCHDOG_H

// Starts watchdog setting time to about `milis` miliseconds.
void watchdog_start(unsigned milis);

// Reloads watchdog.
void watchdog_reload();

#endif // ZADANIE3_WATCHDOG_H
