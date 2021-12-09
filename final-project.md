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
