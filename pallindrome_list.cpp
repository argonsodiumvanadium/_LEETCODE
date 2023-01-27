#include <cstddef>
#include <iostream>
using namespace std;

 struct ListNode {
     int val;
     ListNode *next;
     ListNode() : val(0), next(nullptr) {}
     ListNode(int x) : val(x), next(nullptr) {}
     ListNode(int x, ListNode *next) : val(x), next(next) {}
 };

class Solution {
public:
    bool isPalindrome(ListNode* head) {
        ListNode *fast = head, *slow = head;
        if (head->next == NULL) return true;
        ListNode *reversed_list = new ListNode (slow->val, NULL);
        int len = 1;

        while (fast->next != NULL && fast->next->next != NULL) {
            fast = fast->next->next;
            slow = slow->next;
            reversed_list = new ListNode(slow->val, reversed_list);
            len ++;
        }

        
        cout << "rev " << reversed_list->val << reversed_list->next->val << '\n';
        cout << len << ' ' << len%2 << '\n';

        if (len%2 == 0) {
            cout << "jhel \n";
            reversed_list = reversed_list->next;
        }
        slow = slow->next;
        cout << slow->val << reversed_list->val << '\n';
        cout << slow->val << reversed_list->next->val << '\n';
        if (slow->next == NULL) {
            return slow->val == reversed_list->val;
        }
        while (slow->next != NULL) {
            if (slow->val != reversed_list->val) return false;
            slow = slow->next;
        }

        return true;
    }
};


int main () {
    ListNode* head = new ListNode (1,new ListNode(0,new ListNode(0,NULL)));
    cout << (new Solution())->isPalindrome(head) << '\n';
}
