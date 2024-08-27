inputs: {
  config,
  lib,
  pkgs,
  ...
}: let
  inherit (pkgs.stdenv.hostPlatform) system;
  cfg = config.services.popbobdriver;

  package = inputs.self.packages.${system}.default;
  inherit (lib) mkOption mkEnableOption types mkIf;
in {
  options.services.popbobdriver = {
    enable = mkEnableOption "I FORGOT TO STUDY OKEY?? LIVE ME ALONE";
    package = mkOption {
      type = types.package;
      default = package;
      example = package;
      description = "popbobdriver package";
    };
  };
  config = mkIf cfg.enable {
    systemd.services.popbobdriver = {
      description = "i dont have to explain ";
      wantedBy = ["multi-user.target"];
      wants = ["network.target"];
      after = [
        "network-online.target"
        "NetworkManager.service"
        "systemd-resolved.service"
      ];
      serviceConfig = {
        ExecStart = ''${cfg.package}/bin/popbobdriver'';
        Restart = "always";
      };
    };
  };
}
