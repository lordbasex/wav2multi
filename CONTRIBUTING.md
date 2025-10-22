# 🤝 Contributing to wav2multi

Thank you for your interest in contributing to **wav2multi**! We welcome contributions from the community.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How to Contribute](#how-to-contribute)
- [Contributor License Agreement](#contributor-license-agreement)
- [Development Setup](#development-setup)
- [Coding Guidelines](#coding-guidelines)
- [Submitting Changes](#submitting-changes)
- [Reporting Issues](#reporting-issues)

---

## 📜 Code of Conduct

By participating in this project, you agree to maintain a respectful and professional environment. We expect all contributors to:

- Be respectful and considerate
- Accept constructive criticism gracefully
- Focus on what's best for the project
- Show empathy towards others

---

## 🚀 How to Contribute

There are many ways to contribute:

### 💡 **Suggesting Features**
- Open an issue with the tag `enhancement`
- Describe the feature and its use case
- Explain why it would benefit users

### 🐛 **Reporting Bugs**
- Open an issue with the tag `bug`
- Include steps to reproduce
- Provide system information (OS, Go version, etc.)
- Include error messages and logs

### 📝 **Improving Documentation**
- Fix typos or unclear sections
- Add examples or tutorials
- Translate documentation

### 💻 **Contributing Code**
- Fork the repository
- Create a feature branch
- Make your changes
- Submit a pull request

---

## ⚖️ Contributor License Agreement (CLA)

By submitting a contribution to this project, you agree to the following terms:

### Grant of Copyright License

You grant Federico Pereira (CNSoluciones) a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to:

- Reproduce your contributions
- Prepare derivative works
- Publicly display and perform
- Sublicense and distribute
- Use in commercial products

### Grant of Patent License

You grant Federico Pereira (CNSoluciones) a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable patent license for any patent claims you own that are infringed by your contribution.

### Original Work

You represent that:
- You own the copyright to your contribution, OR
- You have permission from the copyright owner to submit the contribution
- Your contribution does not violate any third-party rights
- You're legally entitled to grant the above licenses

### Licensing

All contributions will be licensed under the **Apache License 2.0**, the same license as the project.

**By submitting a pull request, you acknowledge and agree to these terms.**

---

## 🛠️ Development Setup

### Prerequisites

- **Go 1.23+**
- **CGO enabled** (`CGO_ENABLED=1`)
- **libbcg729** (for G.729 support)
- **Docker** (for testing)
- **Make** (for build automation)

### Setup Steps

```bash
# 1. Fork and clone the repository
git clone https://github.com/lordbasex/wav2multi.git
cd wav2multi

# 2. Install dependencies
go mod download

# 3. Install bcg729 (for local development)
git clone https://github.com/BelledonneCommunications/bcg729
cd bcg729
cmake -S . -B build
cmake --build build --target install
sudo ldconfig

# 4. Build the project
export CGO_ENABLED=1
go build -o transcoding transcoding.go

# 5. Run tests
go test ./...

# 6. Test with Docker
docker build -t wav2multi:test .
docker run --rm wav2multi:test --version
```

---

## 📐 Coding Guidelines

### Go Code Style

Follow the official [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments):

- Use `gofmt` to format code
- Follow idiomatic Go practices
- Write clear, self-documenting code
- Add comments for complex logic
- Keep functions focused and small

### Code Quality

```bash
# Format code
gofmt -w .

# Run linter
golint ./...

# Run vet
go vet ./...

# Run tests
go test -v ./...
```

### Commit Messages

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

**Examples:**
```
feat(codecs): add Opus codec support

Implement Opus encoding and decoding for better quality
at lower bitrates.

Closes #42
```

```
fix(g729): resolve memory leak in encoder

Fixed buffer not being freed after encoding complete.

Fixes #55
```

---

## 📤 Submitting Changes

### Pull Request Process

1. **Fork the repository**
   ```bash
   # On GitHub, click "Fork"
   ```

2. **Create a feature branch**
   ```bash
   git checkout -b feature/my-awesome-feature
   ```

3. **Make your changes**
   - Write clean, documented code
   - Add tests for new features
   - Update documentation

4. **Test your changes**
   ```bash
   go test ./...
   docker build -t wav2multi:test .
   ```

5. **Commit with clear messages**
   ```bash
   git add .
   git commit -m "feat(codecs): add AAC support"
   ```

6. **Push to your fork**
   ```bash
   git push origin feature/my-awesome-feature
   ```

7. **Create Pull Request**
   - Go to the original repository
   - Click "New Pull Request"
   - Select your branch
   - Fill in the PR template

### Pull Request Template

Your PR description should include:

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Tests added/updated
- [ ] All tests pass
- [ ] Tested manually

## Checklist
- [ ] Code follows project style
- [ ] Documentation updated
- [ ] No breaking changes (or documented)
- [ ] License headers added to new files
```

---

## 🐛 Reporting Issues

### Bug Reports

Include:
1. **Clear title** describing the issue
2. **Steps to reproduce**
3. **Expected behavior**
4. **Actual behavior**
5. **Environment details** (OS, Go version, etc.)
6. **Error messages** and logs
7. **Screenshots** if applicable

### Feature Requests

Include:
1. **Clear title** of the feature
2. **Use case** - why is this needed?
3. **Proposed solution** - how should it work?
4. **Alternatives considered**
5. **Additional context**

---

## 🎯 Code Review Process

1. **Automated checks** must pass (linting, tests)
2. **Maintainer review** - typically within 48 hours
3. **Address feedback** - make requested changes
4. **Approval** - once approved, PR will be merged
5. **Attribution** - you'll be credited in release notes

---

## 📄 License

By contributing, you agree that your contributions will be licensed under the **Apache License 2.0**.

---

## 💼 Commercial Contributions

If you're contributing on behalf of a company or want to discuss commercial partnerships:

**Contact:**
- Email: fpereira@cnsoluciones.com
- Company: CNSoluciones
- Website: https://cnsoluciones.com

---

## 🙏 Recognition

Contributors will be:
- Listed in release notes
- Credited in the project README
- Thanked publicly on social media (with permission)

---

## 📞 Questions?

- **General questions**: Open a discussion on GitHub
- **Security issues**: Email fpereira@cnsoluciones.com
- **Commercial inquiries**: Email fpereira@cnsoluciones.com

---

**Thank you for contributing to wav2multi! 🎵**

Made with ❤️ by the wav2multi community

