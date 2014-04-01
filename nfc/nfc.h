#include <nfc/nfc.h>
#include <string.h>
#include <stdlib.h>
int init_globals(char* connstring);
char** get_ids(int* count);
void read_ids(nfc_modulation nm);
char* get_id();
