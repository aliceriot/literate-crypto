typedef struct s_list_node {
	void *data;
	struct s_list_node *next;
} s_list_node;

typedef struct s_list {
	s_list_node *head;
	s_list_node *tail;
} s_list;

/*
 * Singly-linked list
 */
s_list *singly_linked(void *data);
void destroy_s_list(s_list *list);
void print_s_list_node(s_list_node *node);
void print_s_list(s_list *list);
