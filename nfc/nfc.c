#include <nfc/nfc.h>
#include <string.h>
#include <stdlib.h>
#include "nfc.h"

nfc_device *dev;
#define MAX_TARGET_COUNT 16

// NFC Library and device initialization. Should be called precisely once from init()
int init_globals(char* connstring) {
	nfc_context *context;
	nfc_init(&context);
	nfc_connstring nfc_cstring;
	strncpy(nfc_cstring, connstring, NFC_BUFSIZE_CONNSTRING);
	dev = nfc_open(context, nfc_cstring);
	if (dev == NULL) {
		printf("Failed to open NFC Device %s!\n", connstring);
	}
    if (nfc_initiator_init(dev) < 0) {
	    printf("Failed initialization of NFC Device");
    }
    printf("Initialized NFC Device %s.\n", connstring);
}

char** get_ids(int mask, int c) {
	nfc_modulation nm;
	int res = 0;
	nfc_target ant[MAX_TARGET_COUNT];
	char** targets = (char**) malloc(sizeof(char*) * MAX_TARGET_COUNT);
	size_t sz;


    if (mask & 0x1) {
      nm.nmt = NMT_ISO14443A;
      nm.nbr = NBR_106;
      // List ISO14443A targets
      if ((res = nfc_initiator_list_passive_targets(dev, nm, ant, MAX_TARGET_COUNT)) >= 0) {
        int n;
        for (n = 0; n < res; n++) {
		sz = ant[n].nti.nai.szUidLen;
		targets[c] = malloc(sz + 1);
		memcpy(targets[c], ant[n].nti.nai.abtUid, sz);
		targets[c][sz] = 0;
		c++;
        }
      }
    }
    return targets;
}
/*
int foo() {
    if (mask & 0x02) {
      nm.nmt = NMT_FELICA;
      nm.nbr = NBR_212;
      // List Felica tags
      if ((res = nfc_initiator_list_passive_targets(pnd, nm, ant, MAX_TARGET_COUNT)) >= 0) {
        int n;
        if (verbose || (res > 0)) {
          printf("%d Felica (212 kbps) passive target(s) found%s\n", res, (res == 0) ? ".\n" : ":");
        }
        for (n = 0; n < res; n++) {
          print_nfc_target(&ant[n], verbose);
          printf("\n");
        }
      }
    }

    if (mask & 0x04) {
      nm.nmt = NMT_FELICA;
      nm.nbr = NBR_424;
      if ((res = nfc_initiator_list_passive_targets(pnd, nm, ant, MAX_TARGET_COUNT)) >= 0) {
        int n;
        if (verbose || (res > 0)) {
          printf("%d Felica (424 kbps) passive target(s) found%s\n", res, (res == 0) ? ".\n" : ":");
        }
        for (n = 0; n < res; n++) {
          print_nfc_target(&ant[n], verbose);
          printf("\n");
        }
      }
    }

    if (mask & 0x08) {
      nm.nmt = NMT_ISO14443B;
      nm.nbr = NBR_106;

      */
