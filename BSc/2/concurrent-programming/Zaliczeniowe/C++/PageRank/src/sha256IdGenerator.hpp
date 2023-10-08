#ifndef SRC_SHA256IDGENERATOR_HPP_
#define SRC_SHA256IDGENERATOR_HPP_

#include "immutable/idGenerator.hpp"
#include "immutable/pageId.hpp"
#include <array>
#include <fcntl.h>
#include <memory>
#include <unistd.h>

class FileDescriptor {
    int _fd[2];
    bool closed[2] = { false, false };

    void closeHelper(size_t p)
    {
        if (!closed[p]) {
            closed[p] = true;
            if (close(_fd[p])) {
                std::runtime_error("close() error!");
            }
        }
    }

public:
    FileDescriptor()
    {
        if (pipe2(_fd, O_CLOEXEC)) {
            closed[0] = closed[1] = true;
            throw std::runtime_error("pipe() error!");
        }
    }

    int getRead() const
    {
        return _fd[0];
    }

    int getWrite() const
    {
        return _fd[1];
    }

    void closeRead()
    {
        closeHelper(0);
    }

    void closeWrite()
    {
        closeHelper(1);
    }

    ~FileDescriptor()
    {
        closeRead();
        closeWrite();
    }
};

class Sha256IdGenerator : public IdGenerator {
private:
    void set_up_child_pipes(const FileDescriptor& fd_write, const FileDescriptor& fd_read) const
    {
        dup2(fd_write.getRead(), STDIN_FILENO);
        dup2(fd_read.getWrite(), STDOUT_FILENO);
    }

    void child(FileDescriptor& fd_write, FileDescriptor& fd_read) const
    {
        set_up_child_pipes(fd_write, fd_read);

        char* const argv[3] = { (char*)"sha256sum", (char*)"-", NULL };

        execvp("sha256sum", argv);
    }

public:
    virtual PageId generateId(std::string const& content) const
    {

        FileDescriptor fd_write;
        FileDescriptor fd_read;

        switch (fork()) {
        case -1:
            throw std::runtime_error("fork() error!");
        case 0:
            child(fd_write, fd_read);
        }
        fd_write.closeRead();
        fd_read.closeWrite();

        if ((std::string::size_type)write(fd_write.getWrite(), content.c_str(), content.size()) != content.size()) {
            throw std::runtime_error("write() error!");
        }

        fd_write.closeWrite();

        char buf[65];
        if (read(fd_read.getRead(), buf, 64) != 64) {
            throw std::runtime_error("read() error!");
        }
        buf[64] = '\0';

        return PageId(buf);
    }
};

#endif /* SRC_SHA256IDGENERATOR_HPP_ */
