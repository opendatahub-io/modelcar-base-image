.PHONY: Sync
sync: ## Sync dependencies (install from lock file)
	uv sync --all-extras

.PHONY: test
test:
	uv run pytest -x -s -v

.PHONY: build
build:
	uv build

.PHONY: clean
clean:
	rm -rf dist/
	rm -rf build/
	find . -type d -name __pycache__ -exec rm -r {} + 2>/dev/null || true
	find . -type f -name "*.pyc" -delete
	find . -type f -name "*.pyo" -delete
	find . -type d -name "*.egg-info" -exec rm -r {} + 2>/dev/null || true
