// Przerobiony moduł do wypisywania errorów i kończenia działania z laboratoriów.

#ifndef _ERR_
#define _ERR_

/* wypisuje informacje o blednym zakonczeniu funkcji systemowej 
i konczy dzialanie */
void syserr(const char *fmt);

/* wypisuje informacje o bledzie i konczy dzialanie */
void fatal(const char *fmt);

#endif
