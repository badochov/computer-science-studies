#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <errno.h>
#include <string.h>
#include <time.h>

#include "err.h"

#define SIGTOBLOCK (SIGRTMIN + 1)
/*Liczba dzieci*/
#define N   10
#define K   5

 int gotoweDzieci;
  
 int pipe_dsc[2];

void catch (int sig, siginfo_t *info, void *more) { 
   gotoweDzieci ++;
  printf("Rodzic: Dostałem sygnał >>%s<< od %d\n", strsignal(sig), info->si_pid);
  if(gotoweDzieci == K){
        if(close(pipe_dsc[1]))
            syserr("Zamykanie rury nie powiodło się");
  }
        
}

int main ()
{
  
  pid_t parent_pid, child_pid[N];
  int i;
  srand((unsigned) time(0));  // losowość

  parent_pid = getpid();
  struct sigaction action;
  sigset_t s;
  sigemptyset (&s);
  sigaddset(SIGTOBLOCK,&s);
  action.sa_handler = *catch;
  action.sa_mask = s;
  if (sigaction (SIGTOBLOCK, &action, 0) == -1) 
    syserr("sigaction");
  
  if (pipe(pipe_dsc) == -1) syserr("Error in pipe\n");

  for(i=0; i< N; ++i){    //tworzę N dzieci
    int sleep_time = rand()%10 + 1;   // pseudolosowość wymaga losowania w rodzicu
    switch (child_pid[i] = fork()){
      
      case -1:
        syserr("Error in fork\n");
      case 0:
        /*child*/
        printf("Dziecko %d: Zaczynam\n\n", getpid());
        if(close(pipe_dsc[1])){
            syserr("Zamykanie rury nie powiodło się");
        }
        sleep(sleep_time);  // wymagana drzemka
        if(kill(parent_pid, SIGTOBLOCK) < 0)
            syserr("kill");
        char * a = malloc(1);
        if(read(pipe_dsc[0], a,1) == 0){
          if(close(pipe_dsc[0])){
            syserr("Zamykanie rury nie powiodło się");
          }
            printf("Dziecko %d: Kończę\n\n", getpid());
            return 0;
        }
        else{
            syserr("Nieoczekiwana wartość, read");
        }
      default:
        #ifdef D
        printf("Parent %d\n", i);
        #endif
        break;
    }
  }
  /*parent*/
  if(close(pipe_dsc[0])){
      syserr("Zamykanie rury nie powiodło się");
  }
  i = 0;
  while(i<N){
    printf("Czekam na dziecko\n");
    if (wait(0) == -1){
        syserr("wait");
    }else{
      printf("Odebrałem dziecko\n");
      ++i;
    }
  }

  printf("Rodzic: Kończę\n");
  
  return 0;  
}
/**********************************************************************/