docker run --rm -dit \
        --name abc123 \
	--network docker_traefik_proxy \
	-u 0 \
        -l "traefik.enable=true" \
        -l "traefik.http.routers.abc123.rule=PathPrefix(\`/abc123\`)" \
        -l "traefik.http.routers.abc123.entrypoints=web" \
	-l "traefik.http.services.abc123.loadbalancer.server.port=3000" \
	-l "traefik.http.services.abc123.loadbalancer.passHostHeader=true" \
        wetty-docker
