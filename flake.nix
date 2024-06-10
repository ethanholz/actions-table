{
  description = "A simple Go program for converting action.yml to markdown";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };
  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = [
        "x86_64-linux"
        "aarch64-darwin"
      ];
      perSystem = {
        self',
        pkgs,
        ...
      }: {
        packages = {
          default = pkgs.buildGoModule {
            name = "action-table";
            src = ./.;
            vendorHash = "sha256-g+yaVIx4jxpAQ/+WrGKxhVeliYx7nLQe/zsGpxV4Fn4=";
          };
        };
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
