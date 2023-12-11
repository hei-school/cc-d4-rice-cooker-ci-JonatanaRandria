#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>
#include "service.h"

bool plug_in(){
  print_result("Rice cooker pluged");
  return true;
}

bool unplug(){
  print_result("Rice cooker unpluged");
  return false;
}

bool put_something(){
  print_result("Something is there");
  return true;
}

bool empty(){
  print_result("Rice cooker empty");
  return false;
}

bool switch_on(){
  print_result("Rice cooker on");
  return true;
}

bool switch_off(){
  print_result("Rice cooker off");
  return false;
}

void print_error(char msg[]){
  printf("/!\\ ---------------------------------- /!\\\n");
  fprintf(stderr, "\n\t%s \n\n", msg);
  printf("/!\\ ---------------------------------- /!\\\n");
}

void print_result(char msg[]){
  printf(" ---------------------------------- \n");
  printf("\n\t%s \n\n", msg);
  printf(" ---------------------------------- \n");
}