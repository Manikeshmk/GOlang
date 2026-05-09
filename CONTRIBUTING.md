# Contributing to Silent Meeting Summarizer

We welcome contributions from the community! This document provides guidelines for contributing.

## Code of Conduct

Be respectful, inclusive, and professional in all interactions.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/GOlang.git`
3. Create a feature branch: `git checkout -b feature/your-feature`
4. Make your changes
5. Commit with clear messages: `git commit -m "Add feature: description"`
6. Push to your fork: `git push origin feature/your-feature`
7. Open a Pull Request

## Development Setup

```bash
# Install dependencies
make install-deps

# Run tests before submitting
make test
make lint
```

## Code Style

- Follow Go conventions (gofmt)
- Use meaningful variable names
- Add comments for exported functions
- Keep functions focused and testable
- Use interfaces for abstractions

## Testing Requirements

- All new features must include tests
- Aim for 80%+ test coverage
- Use table-driven tests for multiple cases
- Mock external dependencies

## Commit Message Guidelines

- Use present tense: "Add feature" not "Added feature"
- Use imperative mood: "Move cursor to..." not "Moves cursor to..."
- Limit first line to 50 characters
- Reference issues when applicable: "Fixes #123"

## Pull Request Process

1. Update README if needed
2. Update documentation
3. Add/update tests
4. Ensure CI/CD passes
5. Request review from maintainers
6. Address feedback
7. Squash commits if requested

## Reporting Bugs

Use GitHub Issues with:

- Clear title
- Detailed description
- Steps to reproduce
- Expected vs actual behavior
- Go version and OS

## Feature Requests

Include:

- Clear description
- Use cases
- Proposed implementation (if known)
- Examples or mockups

## Performance Considerations

- Profile before optimizing
- Consider memory usage
- Test with realistic data volumes
- Document performance implications

## Documentation

- Update code comments
- Update README if user-facing
- Add examples for new features
- Keep docs in sync with code

## License

By contributing, you agree that your contributions will be licensed under the same MIT License.

---

Thanks for contributing! 🎉
