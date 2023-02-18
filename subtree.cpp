//https://leetcode.com/problems/subtree-of-another-tree/
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
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
	    if (root == NULL || subRoot == NULL) return false;
	    if (root->val == subRoot->val) {
		    if (equal(root, subRoot)) return true;
	    }
	    return isSubtree(root->left,subRoot) || isSubtree(root->right,subRoot);
    }

    bool equal (TreeNode* root, TreeNode* subRoot) {
	    if (root == NULL && subRoot == NULL) return true;
	    if (root == NULL || subRoot == NULL) return false;
	    return root->val == subRoot->val && equal(root->left, subRoot->left) && equal (root->right,subRoot->right);
    }
};
