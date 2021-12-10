# Thực hành làm việc nhóm với Git

Hai bạn cùng đội sẽ vào vai Tom và Jerry, Tom và Jerry sẽ lần lượt thực hiện những nhiệm vụ được giao bên dưới.

Sau khi làm hết các nhiệm vụ hai bạn sẽ đổi vai cho nhau và thực hiện lại một lần nữa.

Kết quả cuối cùng sẽ giống như repo mẫu này: https://github.com/hdnha11/git-cheatsheet.

## Nhiệm vụ 1

- Khởi tạo repository
- Push repository
- Clone repository

### Tom

- Tạo một GitHub repo với tên là `git-cheatsheet`
- Thêm Jerry vào repo và cho quyền ghi
- Tạo một thư mục tên `git-cheatsheet` tại thư mục làm việc trên máy mình

```bash
cd <my-working-directory>
mkdir git-cheatsheet
cd git-cheatsheet
```

- Khởi tạo local git repo

```bash
git init
```

- Thêm file `README.md`

```bash
touch README.md
code README.md # Mở VS Code và copy nội dung bên dưới vào file README.md
```

```markdown
# Git Cheat Sheet

## Git cơ bản

## Huỷ bỏ thay đổi

## Thay đổi lịch sử

## Nhánh

## Remote Repo

## Cấu hình Git

## Git Log

## Git Diff

## Git Reset

## Git Rebase

## Git Pull

## Git Push
```

- Tạo commit đầu tiên

```bash
git add README.md
git status
git commit -m "First commit"
git log
```

- Push local repo lên remote repo

```bash
git remote add origin git@github.com:<your-username>/git-cheatsheet.git
git branch -M main
git push -u origin main
```

### Jerry

- Clone `git-cheatsheet` từ Tom

```bash
cd <my-working-directory>
git clone git@github.com:<tom-username>/git-cheatsheet.git
cd git-cheatsheet
git log
```

## Nhiệm vụ 2

- Tạo nhánh và làm việc trên nhánh
- Tạo Pull Request
- Review Pull Request
- Merge Pull Request

### Tom

- Tạo nhánh feature

```bash
git checkout -b feature/task-2-tom
git log
# bạn sẽ thấy là hiện tại HEAD đang trỏ cùng lúc tới nhánh `main` và `feature/task-2-tom`
```

- Thêm nội dung cho các phần
	- Git cơ bản
	- Huỷ bỏ thay đổi

Dùng text editor mở file `README.md` và update nội dung như phần bên dưới và lưu lại.

```markdown
## Git cơ bản

| Lệnh                            | Mô tả                                                                                                    |
| ------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `git init <directory>`          | Tạo repo trống trong thư mục được chỉ định. Nếu không cung cấp thư mục sẽ tạo repo cho thư mục hiện tại. |
| `git clone <repo>`              | Clone remote repo có đường dẫn `<repo>` về máy.                                                          |
| `git config user.name <name>`   | Đặt lại tên tác giả.                                                                                     |
| `git add <directory> or <file>` | Đưa thư mục hoặc file vào index để chuẩn bị cho commit tiếp theo.                                        |
| `git commit -m "<message>"`     | Commit tất cả file trong index với mô tả là `<message>`.                                                 |
| `git status`                    | Xem trạng thái của các file.                                                                             |
| `git log`                       | Hiện toàn bộ lịch sử commit.                                                                             |
| `git diff`                      | Hiển thị những thay đổi hiện tại mà bạn tạo ra so với index (chưa được add vào index).                   |

## Huỷ bỏ thay đổi

| Lệnh                  | Mô tả                                                            |
| --------------------- | ---------------------------------------------------------------- |
| `git revert <commit>` | Tạo một commit đảo ngược với `<commit>` để huỷ bỏ `<commit>`.    |
| `git reset <file>`    | Bỏ `<file>` ra khỏi index (vẫn giử lại nội dung thay đổi).       |
| `git clean -df`       | Xoá bỏ tất cả file và thư mục không muốn đưa vào Git (file rác). |
```

- Kiểm tra thay đổi với Git

```bash
git diff
```

- Commit thay đổi

```bash
git add .
git status
git commit -m 'Update Git cơ bản & Huỷ bỏ thay đổi'
git log
```

- Push lên nhánh remote

```bash
git push origin -u feature/task-2-tom
```

- Tạo pull request trên GitHub và nhờ Jerry review
- Review và approve pull request của Jerry
- Chờ Jerry merge pull request
- Đợi Jerry approve pull request và nhấn "Merge pull request" để merge PR của mình
- Chuyển về nhánh `main` và pull code mới về (luôn pull code mới về trước khi làm việc mới)

```bash
git checkout main
git pull
git log
```

Dùng text editor để mở file `README.md` bạn sẽ thấy file mới đã bao gồm cả 2 phần thay đổi của cả 2 bạn Tom & Jerry.

Chúc mừng Tom đã hoàn thành nhiệm vụ số 2.

### Jerry

- Tạo nhánh feature

```bash
git checkout -b feature/task-2-jerry
git log
# bạn sẽ thấy là hiện tại HEAD đang trỏ cùng lúc tới nhánh `main` và `feature/task-2-jerry`
```

- Thêm nội dung cho các phần
	- Thay đổi lịch sử
	- Nhánh
	- Remote Repo

Dùng text editor mở file `README.md` và update nội dung như phần bên dưới và lưu lại.

```markdown
## Thay đổi lịch sử

| Lệnh                 | Mô tả                                                                                  |
| -------------------- | -------------------------------------------------------------------------------------- |
| `git commit --amend` | Commit nối vào commit gần nhất (không tạo commit mới).                                 |
| `git rebase <base>`  | Dùng `<base>` làm base mới. (`<base>` có thể là nhánh, tag, commit).                   |
| `git reflog`         | Xem lịch sử thay đổi của `HEAD` (rất hữu ích khi lỡ tay làm mất code và muốn tìm lại). |

## Nhánh

| Lệnh                       | Mô tả                                                            |
|----------------------------|------------------------------------------------------------------|
| `git branch`               | Liệt kê tất cả nhánh trong repo. Thêm `-a` để xem remote branch. |
| `git branch <branch>`      | Tạo nhánh mới với tên `<branch>`.                                |
| `git checkout <branch>`    | Chuyển sang nhánh có tên `<branch>`.                             |
| `git checkout -b <branch>` | Tạo và chuyển sang nhánh `<branch>`.                             |
| `git branch -D <branch>`   | Xoá nhánh `<branch>`.                                            |
| `git merge <branch>`       | Merge nhánh `<branch>` vào nhánh hiện tại.                       |

## Remote Repo

| Lệnh                          | Mô tả                                                                                               |
|-------------------------------|-----------------------------------------------------------------------------------------------------|
| `git remote add <name> <url>` | Tạo kết nối với remote repo tại `<url>` (sau đó có thể dùng `<name>` như là shortcut).              |
| `git fetch <branch>`          | Fetch nhánh `<branch>` từ remote repo. Không cung cấp tên nhánh có nghĩa là fetch tất cả các nhánh. |
| `git pull`                    | Fetch nhánh hiện tại và đồng thời merge nhánh remote vào local.                                     |
| `git push <remote> <branch>`  | Push nhánh lên remote repo.                                                                         |
```

- Kiểm tra thay đổi với Git

```bash
git diff
```

- Commit thay đổi

```bash
git add .
git status
git commit -m 'Update Thay đổi lịch sử & Nhánh & Remote Repo'
git log
```

- Push lên nhánh remote

```bash
git push origin -u feature/task-2-jerry
```

- Tạo pull request trên GitHub và nhờ Tom review
- Review và approve pull request của Tom
- Đợi Tom approve pull request và nhấn "Merge pull request" để merge PR của mình trước Tom
- Đợi Tom merge PR của Tom
- Chuyển về nhánh `main` và pull code mới về (luôn pull code mới về trước khi làm việc mới)

```bash
git checkout main
git pull
git log
```

Dùng text editor để mở file `README.md` bạn sẽ thấy file mới đã bao gồm cả 2 phần thay đổi của cả 2 bạn Tom & Jerry.

Chúc mừng Jerry đã hoàn thành nhiệm vụ số 2.

## Nhiệm vụ 3

- Giải quyết đụng độ (conflict)

### Tom

- Tạo nhánh feature

```bash
git checkout main
git pull
git checkout -b feature/task-3-tom
```

- Thêm nội dung cho các phần
	- Cấu hình Git

Dùng text editor mở file `README.md` và update nội dung như phần bên dưới và lưu lại.

```markdown
## Cấu hình Git

| Lệnh                                       | Mô tả                                                        |
|--------------------------------------------|--------------------------------------------------------------|
| `git config --global user.name <name>`     | Đặt tên tác giả cho tất cả commit của user hiện tại.         |
| `git config --global user.email <email>`   | Đặt email tác giả cho tất cả commit của user hiện tại.       |
```

- Commit và push

```bash
git diff
git add .
git status
git commit -m 'Tom Update Cấu hình Git'
git log
git push origin -u feature/task-3-tom
```

- Tạo pull request trên GitHub và chờ Jerry review
- Review và approve pull request của Jerry

Sau khi Jerry merge pull request của bạn ấy thì Tom không thể merge pull request của mình vì đã xảy ra conflict.

- Xử lý đụng độ

```bash
git fetch
git rebase origin/main
# Git sẽ thông báo danh sách các file đang bị conflict:
#   Auto-merging README.md
#   CONFLICT (content): Merge conflict in README.md
#   error: could not apply fc18839... Tom Update Cấu hình Git
#   hint: Resolve all conflicts manually, mark them as resolved with
#   hint: "git add/rm <conflicted_files>", then run "git rebase --continue".
#   hint: You can instead skip this commit: run "git rebase --skip".
#   hint: To abort and get back to the state before "git rebase", run "git rebase --abort".
# Mở text editor và xử lý tất cả conflict (merge code với nhau) (ở đây là file README.md)
# Ở đây do mình muốn giử lại cả phần của Tom và Jerry nên chỉ đơn giản là xoá 3 dòng bắt đầu với (<<<<<<<, =======, >>>>>>>) đi là được
# Sau khi xử lý tất cả conflict nhớ add file vào index (staging area)
git add .
git rebase --continue
git push -f # Force push bởi vì chúng ta đã thay đổi lịch sử của nhánh
```

- Quay lại GitHub để review lại PR
- Merge PR
- Chuyển về nhánh `main` và pull code mới về

```bash
git checkout main
git pull
git log
```

Chúc mừng Tom đã thực hiện xong nhiệm vụ 3.

### Jerry

- Tạo nhánh feature

```bash
git checkout main
git pull
git checkout -b feature/task-3-jerry
```

- Thêm nội dung cho các phần
	- Cấu hình Git

Dùng text editor mở file `README.md` và update nội dung như phần bên dưới và lưu lại.

```markdown
## Cấu hình Git

| Lệnh                                       | Mô tả                                                        |
|--------------------------------------------|--------------------------------------------------------------|
| `git config --system core.editor <editor>` | Đặt text editor mong muốn (mặc định là Vim) cho tất cả user. |
```

- Commit và push

```bash
git diff
git add .
git status
git commit -m 'Jerry Update Cấu hình Git'
git log
git push origin -u feature/task-3-jerry
```

- Tạo pull request trên GitHub và chờ Tom review
- Merge PR sau khi Tom approve
- Review và approve pull request của Tom
- Chuyển về nhánh `main` và pull code mới về

```bash
git checkout main
git pull
git log
```

Chúc mừng Jerry đã thực hiện xong nhiệm vụ 3.

## Nhiệm vụ 4

- Hoàn thành dự án

Ở nhiệm vụ này 2 bạn Tom & Jerry tự do trao đổi và phân công phần công việc còn lại để hoàn thành dự án.

```markdown
## Git Log

| Lệnh                           | Mô tả                                      |
|--------------------------------|--------------------------------------------|
| `git log --stat`               | Hiện thị thêm thông tin file được sửa đổi. |
| `git log --author="<pattern>"` | Lọc commit theo tác giả.                   |
| `git log --grep="<pattern>"`   | Lọc commit theo commit message.            |
| `git log <since>..<until>`     | Hiện thị commit trong khoảng.              |
| `git log -- <file>`            | Hiện thị commit của một file cụ thể.       |
| `git log --graph`              | Hiện thị lịch sử dưới dạng đồ thị.         |

## Git Diff

| Lệnh                | Mô tả                                           |
|---------------------|-------------------------------------------------|
| `git diff HEAD`     | So sánh working directory với commit cuối cùng. |
| `git diff --cached` | So sánh index với commit cuối cùng.             |

## Git Reset

| Lệnh               | Mô tả                                                                              |
|--------------------|------------------------------------------------------------------------------------|
| `git reset`        | Reset khu vực index (staging area) giống với commit gần nhất.                      |
| `git reset --hard` | Reset khu vực index (staging area) và working directory giống với commit gần nhất. |

## Git Rebase

| Lệnh                   | Mô tả                                                                  |
|------------------------|------------------------------------------------------------------------|
| `git rebase -i <base>` | Rebase với chế độ "tương tác", Git sẽ mở editor cho chúng ta thao tác. |

## Git Pull

| Lệnh                         | Mô tả                                 |
|------------------------------|---------------------------------------|
| `git pull --rebase <remote>` | Pull với chế độ rebase thay vì merge. |

## Git Push

| Lệnh                        | Mô tả                                                 |
|-----------------------------|-------------------------------------------------------|
| `git push <remote> --force` | Force push, sẽ ghi đè nhánh remote. Cẩn thận khi làm. |
| `git push <remote> --tags`  | Push local tag lên remote.                            |
```
