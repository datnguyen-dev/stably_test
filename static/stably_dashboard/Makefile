dev:
	npm run serve

release: clean
	npm run build

clean:
	rm -rf dist

docker: release
	docker build -t stably_dashboard:latest .