---
author: Nha Hoang
date: MMMM DD, YYYY
---

# Git cơ bản và nâng cao

- Phần I: Bắt đầu với Git
- Phần II: Làm việc nhóm với Git

---

## Git dành cho người mới bắt đầu

- Giới thiệu về Git
- Linux/Unix Command Line
- Cách khởi tạo một Git repository
- Cách tạo commit
- Cách tạo một remote repository với GitHub, GitLab, Bitbucket
- Cách đồng bộ giữa local và remote repository
- Cách chia sẽ repository của mình với người khác
- Cách clone một repository được chia sẽ từ người khác
- Git GUI
- GUI vs CLI
- Thực hành tạo trang GitHub profile cá nhân

---

### Giới thiệu về Git

- Quản lý version cho tài liệu (Plain Text)
- Mã nguồn Linux kernel
- Hệ thống phân tán
- GitHub vs GitLab vs Bitbucket

---

### Cài đặt Git

#### macOS

```bash
brew install git
```

#### Ubuntu

```bash
sudo apt-get install git
```

#### Windows

https://git-scm.com/download/win

#### Kiểm tra cài đặt

```bash
git --version
```

---

### Linux Command Line căn bản

- Command Line Interface (CLI)
- Cách mở Terminal
- Xem thư mục hiện tại `pwd`
- Liệt kê nội dung thư mục `ls -l <dir_path>`
- Hiện nội dung ẩn `ls -la <dir_path>`
- Một số thư mục đặc biệt `/`, `~`, `.`, `..`
- Chuyển thư mục `cd <dir_path>`
- Tạo thư mục `mkdir <dir_path>`
- Xoá thư mục `rmdir <dir_path>` (`rm -r <dir_path>`)
- Tạo file `touch <file_path>`
- Xoá file `rm <file_path>`
- Sửa file `nano <file_path>`, `vi <file_path>`
- Xem nội dung file `cat <file_path>`
- Chạy lệnh dưới quyền super user `sudo <cmd>`
- Trợ giúp `<cmd> -h`, `man <cmd>`

#### Mẹo

- Dùng TAB để được nhắc code (autocomplete)
- Dùng phím mũi tên lên/xuống để xem lại các lệnh gần đây
- Ctrl+R để tìm kiếm một lệnh đã chạy trong quá khứ
- Ctrl+C để dừng lệnh đang chạy

---

### Khởi tạo Git repository

- Là một thư mục bất kỳ
- Chứa thư mục ẩn `.git`
- `git init`

---

### Thêm file vào Git

- Xem trạng thái của repository `git status`
- Xem nội dung thay đổi `git diff`
- File mới thêm sẽ có trạng thái *Untracked*
- Working tree
- Staging area
- `git add`

![Vòng đời của trạng thái file](https://git-scm.com/book/en/v2/images/lifecycle.png)

---

### .gitignore

- Loại trừ một số file khỏi Git
- *.DS_Store*, log file...
- Git sẽ bỏ qua danh sách file trong file ẩn `.gitignore`
- Có thể dùng ký tự wildcard (`*`) với tên file (`*.log`...)

---

### Tạo commit

---

### Huỷ bỏ thay đổi

- `git reset`
- `git checkout <path>`
- `git clean -df`

---

### Cấu hình Git

- User name
- User email

---

### Tạo remote repository

- GitHub
- GitLab
- Bitbucket

---

### Đẩy nội dung local lên remote repository

> https://onlywei.github.io/explain-git-with-d3/#push

---

### Thêm thành viên vào GitHub repository

- Public vs Private repositories

---

### Kéo nội dung mới từ remote repository

- `git fecth`
- `git pull`

> https://onlywei.github.io/explain-git-with-d3/#fetch
> https://onlywei.github.io/explain-git-with-d3/#pull

---

### Quy trình làm việc điển hình

---

### Một số công cụ GUI cho Git

- [Sourcetree](https://www.sourcetreeapp.com)
- Visual Studio Code (integrated SCM)

> SCM: Source Control Management

---

### GUI vs CLI

Lý do nên học Git CLI trước:

- Giống nhau ở mọi nơi
- Đầy đủ tính năng
- Dễ dàng tìm kiếm sự trợ giúp

Lý do dùng Git GUI:

- Xem lịch sử dễ dàng hơn
- Xem nội dung thay đổi (side-by-side)

---

### Tạo GitHub Profile Repository

- Repository đặc biệt (trùng tên với username)
- Markdown
- Cách căn chỉnh hình ảnh với GitHub markdown
- Tự động sinh thông tin stats

> https://github.com/anuraghazra/github-readme-stats

---

## Làm việc nhóm với Git

- Giới thiệu về mô hình làm việc nhóm với Git
- Cách tạo và quản lý nhánh
- Merge và rebase nhánh
- Cách thay đổi lịch sử của nhánh
- Đụng độ là gì
- Cách xử lý khi xảy ra đụng độ
- Giới thiệu về Git Flow
- Pull Request là gì
- Review và merge Pull Request
- CI/CD với GitHub Actions
- Thực hành làm việc nhóm với Git

---

### Mô hình làm việc nhóm phân tán

---

### Tạo nhánh và quản lý nhánh

- Nhánh là gì?
- Tạo nhánh `git branch <branch_name>`
- Tạo nhánh và checkout `git checkout -b <branch_name>`
- Liệt kê nhánh `git branch -a`
- Chuyển nhánh `git switch <branch_name>`, `git checkout <branch_name>`
- Xoá nhánh `git branch -d <branch_name>`
- Xoá nhánh force `git branch -D <branch_name>`
- Commit vào nhánh
- Push/Pull nhánh

> https://onlywei.github.io/explain-git-with-d3/#branch

---

### Merge nhánh

> https://onlywei.github.io/explain-git-with-d3/#merge

---

### Rebase nhánh

> https://onlywei.github.io/explain-git-with-d3/#rebase

---

### Thay đổi lịch sử

- Xoá commit
- Thay đổi commit message
- Thay đổi tác giả commit
- Gộp nhiều commit thành một

> https://onlywei.github.io/explain-git-with-d3/#reset

---

### Revert commit

- Revert commit thường
- Revert merge commit (có nhiều hơn 1 cha)

> https://onlywei.github.io/explain-git-with-d3/#revert

---

### Đụng độ

- Merge/Rebase nhánh
- Cùng sửa trên cùng dòng của cùng file
- Sửa trên file đã bị xoá

---

### Giải quyết đụng độ

- Xem thay đổi
- Merge code
- Đánh dấu đã giải quyết xong
- Tiếp tục quá trình merge

---

### Giới thiệu về Git Flow

---

### Pull Request là gì

---

### Review và Merge Pull Request

---

### CI/CD với GitHub Actions

---

### Thực hành làm việc nhóm với Git
