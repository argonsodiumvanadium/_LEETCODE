#include <stdio.h>
#include <stdlib.h>

typedef struct node Node;
struct node {
    int   value;
    Node *right;
    Node *left;
}

Node* nalloc (int);
Node* construct_tree (Node*, int*, int);
void  print_tree(Node*);
int   ipow(int, int);

Node*
nalloc (int val) {
    Node* node = malloc(sizeof(Node));
    node->right = NULL;
    node->left  = NULL;
    node->value = val;

    return node;
}

Node*
construct_tree (Node* root,int* values, int level) {
    if (!*values) return root;
    root->value = *values;
    construct_tree(root->left,level+values,ipow(2,level));
    construct_tree(root->right,level+values,ipow(2,level));

    return root;
}

void
print_tree (Node* root) {
    if (!root->value) return;
    printf("%d ",root->value);
    print_tree(root->right);
    print_tree(root->left);
}

int
ipow (int x, int n) {
    return n?x*ipow(x,--n):1;
}

int
main () {
    int values[7] = {1,2,3,4,5,6,0};
    Node* root = nalloc(0);
    construct_tree(root,values,1);
    print_tree(root);
    putchar('\n');
}
