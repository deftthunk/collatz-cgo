#include <stdio.h>
#include "collatz.h"

unsigned long collatz(unsigned long start, unsigned long blockSize, unsigned long ret) {
	unsigned long localRec = 0;
	unsigned long localWin = 0;
	unsigned long orig = start;
	unsigned long local = 0;
	unsigned long chain = 0;

	while(start < (orig + blockSize)) {		
		chain = 0;
		local = start;

		while(1) {
			++chain;
		
			if(!(local & 1)) {
				local /= 2;
				continue;
			}
			else if(local == 1) {
				break;
			}
			else {
				local = local * 3 + 1;
				continue;
			}
		}

		if(chain > localRec) {
			localRec = chain;
			localWin = start;
		}
		++start;
	}

	if(localRec > ret) {
		printf("\r                                              ");
		printf("\rrecord: %lu\t%lu\n", localRec, localWin);
		fflush(stdout);
		ret = localRec;
	}

	return ret;
}
