//https://leetcode.com/problems/maximum-depth-of-binary-tree/submissions/
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int maxDepth(TreeNode* root) {
        if (root == NULL) return 0;
        if (root->right == NULL && root->left == NULL) return 1;
        int left = maxDepth(root->left);
        int right = maxDepth(root->right);
        int min = left < right ? right:left;
        min = left == 0 ? right : min;
        min = right == 0? left : min;
        return 1 +  min;
    }
};
