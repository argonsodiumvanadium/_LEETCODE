//https://leetcode.com/problems/missing-number/

int missingNumber(int* itr, int size){
    short int key=size;
    for (short int i=0;i<size;i++,itr++) {
        key = key ^ *itr ^ i;
    }
    return key;

}
