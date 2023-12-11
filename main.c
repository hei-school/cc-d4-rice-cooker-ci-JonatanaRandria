#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>
#include "./module/service.h"

bool status = true;
int choice;

bool is_pluged_in = false;
bool is_switched_on = false;
bool is_there_anything = false;
bool is_cooking = false;

void main(){
  while(status){
    print_result("RICE COOKER");
    printf("\nMenu:\n\n");
    printf("1- Plug In\n");
    printf("2- Put Anything\n");
    printf("3- Switch On\n");
    printf("4- Switch Off\n");
    printf("5- Empty\n");
    printf("6- Unplug\n");
    printf("7- Exit\n\n");
    printf("I want to : ");
    scanf("%d", &choice);
    switch(choice){
      case 1:
        if(is_pluged_in){
          print_error("Rice Cooker is already pluged!");
        } else {
          is_pluged_in = plug_in();
        }
        break;
      
      case 2:
        if(is_there_anything){
          print_error("Something is already there !");
        } else {
          is_there_anything = put_something();
        }
        break;
      
      case 3:
        if(is_there_anything && is_pluged_in){
          is_cooking = switch_on();
          is_switched_on = is_cooking;
        } else {
          print_error("switch on the power or this is empty,pouring more rice !");
        }
        break;
      
      case 4:
        if(is_cooking && is_switched_on){
          is_cooking = switch_off();
          is_switched_on = is_cooking;
        } else {
          print_error("Rice Cooker is already off !");
        }
        break;
      
      case 5:
        if(!is_there_anything){
          print_error("Something is already empty !");
        } else if (is_cooking) {
          print_error("The cooker is already in cooking process !");
        } else {
          is_there_anything = empty();
        }
        break;
      
      case 6:
        if(!is_pluged_in){
          print_error("Cooker is already unpluged !");
        } else if (is_switched_on || is_cooking) {
          print_error("Cooker is already switched on !");
        } else {
          is_pluged_in = unplug();
        }
        break;

      case 7:
        print_result("Thanks,have a nice and Enjoy your meal !");
        status = false;
        break;
      
      default:
        printf("");
        break;
    }
  }
}