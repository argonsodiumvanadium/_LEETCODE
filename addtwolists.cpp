// https://leetcode.com/problems/add-two-numbers/
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode () : val(0), next(nullptr) {}
    ListNode (int x) : val(x), next(nullptr) {}
    ListNode (int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
    public:
        ListNode* addTwoNumbers (ListNode *l1, ListNode* l2, int carry=0) {
            auto bnull = l1 == nullptr && l2 == nullptr;
            l1 = l1 == nullptr ? new ListNode(0) : l1;
            l2 = l2 == nullptr ? new ListNode(0) : l2;

            auto sum = l1->val + l2->val + carry;
            auto rmn = sum%10;
            carry = sum/10;

            if (!sum && bnull) return nullptr;
            ListNode *node = new ListNode(rmn);
            node->next = addTwoNumbers(l1->next, l2->next, carry);

            return node;
        }
};
