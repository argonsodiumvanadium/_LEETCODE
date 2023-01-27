/**
 * https://leetcode.com/problems/average-of-levels-in-binary-tree/
 */
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
}

double*
averageOfLevels (struct TreeNode* root, int* returnSize) {
    *returnSize = 1;
    struct TreeNode *stack, *itr;
    itr = root;

    while (itr->right)
}
