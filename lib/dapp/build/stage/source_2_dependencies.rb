module Dapp
  module Build
    module Stage
      # Source2Dependencies
      class Source2Dependencies < SourceDependenciesBase
        def initialize(application, next_stage)
          @prev_stage = Artifact.new(application, self)
          super
        end

        protected

        def dependencies
          application.builder.infra_setup_checksum
        end
      end # Source2Dependencies
    end # Stage
  end # Build
end # Dapp
