#include <stdio.h>
#include <sys/ioctl.h>
#include <sys/socket.h>
#include <linux/vm_sockets.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <errno.h>
#include <unistd.h>

int main(int argc, char** argv) {
    int fd;
    int err;
    fd = open("/dev/vsock", O_RDONLY);
    if (fd == -1) {
        err = errno;
        printf("open() failed with %d\n", err);
        return 1;
    }
    unsigned int cid = 0;
    err = ioctl(fd, IOCTL_VM_SOCKETS_GET_LOCAL_CID, &cid);
    if (err == -1) {
        printf("ioctl failed with %d\n", errno);
        close(fd);
        return 1;
    }
    printf("cid=%d\n", cid);
    close(fd);
    return 0;
}
