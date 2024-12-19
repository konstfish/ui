deploy:
	KO_DOCKER_REPO=ghcr.io/konstfish/ui ko build --bare --platform=all
	kn service update konstfish-ui --annotation deploy-timestamp=$(date +%s)