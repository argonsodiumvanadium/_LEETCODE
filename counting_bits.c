// https://leetcode.com/problems/counting-bits/submissions/

int *countingBits (int n, int* returnSize) {
    returnSize = (int*) malloc(sizeof(int)*n);
    *returnSize = 0;
    *(returnSize+1) = 1;

    for (int i=2;i<n;i++) {
        *(returnSize+i)= i&1 + *(returnSize+i>>1)
    }

    return returnSize;
}
