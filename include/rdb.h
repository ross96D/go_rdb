#include <stdint.h>

struct Result
{
    void* database;
    char* error;
};

struct Bytes
{
    char* ptr;
    uint64_t len;
};


struct Result create(char* path);
struct Bytes search(void* db, char* key);
int insert(void* db, char* key, struct Bytes value);
int update(void* db, char* key, struct Bytes value);
int delete(void* db, char* key);
