
echo "# epoccipital" | tee README.md
echo "" | tee -a README.md
epoccipital --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## apikeys" | tee -a README.md
echo "" | tee -a README.md
epoccipital  apikeys  --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## completion" | tee -a README.md
echo "" | tee -a README.md
epoccipital  completion  --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## debug" | tee -a README.md
echo "" | tee -a README.md
epoccipital  debug  --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## generate" | tee -a README.md
echo "" | tee -a README.md
epoccipital  generate  --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## hiddenserve" | tee -a README.md
echo "" | tee -a README.md
epoccipital  hiddenserve  --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## mockoidc" | tee -a README.md
echo "" | tee -a README.md
epoccipital  mockoidc  --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## namespaces" | tee -a README.md
echo "" | tee -a README.md
epoccipital  namespaces  --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## nodes" | tee -a README.md
echo "" | tee -a README.md
epoccipital  nodes  --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## preauthkeys" | tee -a README.md
echo "" | tee -a README.md
epoccipital  preauthkeys --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## routes" | tee -a README.md
echo "" | tee -a README.md
epoccipital  routes --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md
echo "## version" | tee -a README.md
echo "" | tee -a README.md
epoccipital  version --help | sed 's|  |        |g' | tr ':' '\n' | tee -a README.md
echo "" | tee -a README.md