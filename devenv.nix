{ pkgs, config, ... }:

let
  # Dev port (8080 is commonly taken locally); web.go defaults to 8080 for
  # Fly/Docker when PORT is unset.
  port = 3000;
in
{
  env.PORT = toString port;

  # https://devenv.sh/languages/
  # Provides go, gopls, delve and friends.
  languages.go.enable = true;

  # https://devenv.sh/packages/
  packages = [
    pkgs.air # hot reload
    pkgs.flyctl # Fly.io CLI (fly/flyctl)
  ];

  # https://devenv.sh/processes/
  # `devenv up`: rebuild and restart on Go, template, content, or asset changes
  # (views are embedded and responses are cached in memory). See .air.toml.
  processes.web = {
    exec = "cd ${config.devenv.root} && exec air";
    ready.http.get = {
      inherit port;
      path = "/";
    };
  };

  # https://devenv.sh/scripts/
  # Compile locally first to fail fast, then let Fly build the Dockerfile
  # remotely. Extra args pass through, e.g. `deploy --local-only`.
  scripts.deploy.exec = ''
    set -euo pipefail
    cd "${config.devenv.root}"
    go build -o /dev/null .
    exec flyctl deploy "$@"
  '';
}
