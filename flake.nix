{
  description = "A simple Go program for converting action.yml to markdown";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };
  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "aarch64-darwin"
      ];
      perSystem = {
        self',
        pkgs,
        system,
        ...
      }: let
        commonArgs = rec {
            version = "v0.1.1";
            name = "action-table";
            pname = "action-table";
            src = ./.;
            vendorHash = "sha256-g+yaVIx4jxpAQ/+WrGKxhVeliYx7nLQe/zsGpxV4Fn4=";
            ldflags = ["-X main.Version=${version}"];
            CGO_ENABLED = "0";
            doCheck = false;
        };
        package = pkgs.buildGoModule(commonArgs // {});
        ci = pkgs.buildGoModule(commonArgs // {
            installPhase = ''
                  runHook preInstall

                  mkdir -p $out
                  initial="$GOPATH/bin/action-table"
                  final="$GOPATH/bin/action-table-${system}"
                  dir="$GOPATH/bin"
                  [ -e "$initial" ] && mv $initial $final
                  dir="$GOPATH/bin"
                  [ -e "$dir" ] && cp -r $dir $out

                  runHook postInstall
            '';
        });
        in{
        packages = {
            default = package; 
            ci = ci;
        };
        checks.default = package;
        devShells = {
          default = pkgs.mkShell {
            inputsFrom = [self'.packages.default];
          };
        };
        formatter = pkgs.alejandra;
      };
    };
    nixConfig = {
        extra-substituters = ["https://action-table.cachix.org"];
        extra-trusted-public-keys = ["action-table.cachix.org-1:IbI8XIJqLPAuAPS4c9X86ZJ0vgwwJpZHXO38IbknRAQ="];
    };
}
