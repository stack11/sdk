#!/bin/bash

set -Eeuxo pipefail
cd "$( dirname "${BASH_SOURCE[0]}" )/.."

source scripts/prep.sh

to_git() {
  lang=$1
  gitdir="repos/${GIT_REPO}-${lang}"
  srcdir="clients/${PROJECT}/${lang}"

  mkdir -p ${gitdir} || true
  git clone "git@github.com:ory/${GIT_REPO}-${lang}.git" "${gitdir}" || true

  (cd "${gitdir}"; git fetch origin || true; git checkout master || true; git reset --hard HEAD || true; git pull -ff || true; git checkout -b "release-$(date +%s)" master)

  # rm -rf "${srcdir:?}/*"
  rm -rf "${gitdir:?}/*"
  cp -R "${srcdir}/." "${gitdir}"
  ls -la "${gitdir}"

  (cd "${gitdir}"; git add -A || true; (git commit -a  -F- <<EOF
chore: Regenerate OpenAPI client ${VERSION}

Version: ${VERSION}
EOF
) || true)

  if [ "${2}" = "yes" ]; then
        (cd "${gitdir}"; git tag -a "${VERSION}" -m "${VERSION}")
  else
        # empty, do nothing!
        echo "not tagging"
  fi

  (cd "${gitdir}"; git push origin --tags HEAD:master)
}

upstream() {
  git add -A
  git commit --allow-empty -a -m "chore: regenerate OpenAPI client ${VERSION}"
  git push origin
}

typescript() {
  dir="clients/${PROJECT}/typescript"

  (cd "${dir}"; npm install; npm run build)
  (cd "${dir}"; npm version -f --no-git-tag-version "${VERSION}" || true; npm publish --access public)
}

java() {
  to_git "java" "no"

  gitdir="repos/${GIT_REPO}-java"
  (cd "${gitdir}"; mvn clean)

  version=$(echo "${VERSION}" | sed "s/^v//")

  (cd "${gitdir}"; mvn release:prepare \
    -Dresume=false \
    -DreleaseVersion="${version}" \
    -Dtag="${VERSION}" \
    -DdevelopmentVersion="${version}-SNAPSHOT" \
    -Darguments="-Dgpg.passphrase=${MVN_PGP_PASSPHRASE} -Dgpg.keyname=${MVN_PGP_KEYNAME}")

  (cd "${gitdir}"; mvn release:perform)
  (cd "${gitdir}"; git push origin --tags HEAD:master)
}

php() {
  dir="clients/${PROJECT}/php"

  (cd "${dir}"; composer install)
  to_git "php" "yes"
}

ruby() {
  dir="clients/${PROJECT}/ruby"

  gemspec="ory-${PROJECT}-client.gemspec"
  gemfile="ory-${PROJECT}-client-${GEM_VERSION}.gem"
  if [ ${PROJECT} == "client" ]; then
    gemspec="ory-client.gemspec"
    gemfile="ory-client-${GEM_VERSION}.gem"
  fi

  (cd "${dir}"; rm *.gem || true; gem build "${gemspec}"; gem push "${gemfile}")
}

golang() {
  dir="clients/${PROJECT}/go"

  (cd "${dir}"; go mod tidy)
  to_git "go" "yes"
}

python() {
  dir="clients/${PROJECT}/python"
  (cd "${dir}"; rm -rf "dist" || true; python3 setup.py sdist bdist_wheel; python3 -m twine upload "dist/*")
}

dotnet() {
  dir="clients/${PROJECT}/dotnet"

  (cd "${dir}"; VERSION=${RAW_VERSION} command dotnet pack -o .)

  nupkg_name="Ory.${PROJECT_UCF}.Client.${RAW_VERSION}.nupkg"
  if [ ${PROJECT} == "client" ]; then
    nupkg_name="Ory.Client.${RAW_VERSION}.nupkg"
  fi

  (cd "${dir}"; command dotnet nuget push "${nupkg_name}" \
  --api-key "${NUGET_API_KEY}" \
  --source https://api.nuget.org/v3/index.json)
}

dart() {
  dir="clients/${PROJECT}/dart"

  mkdir -p ~/.pub-cache || true
  set -x
  cat <<EOF > ~/.pub-cache/credentials.json
{
  "accessToken":"${DART_ACCESS_TOKEN}",
  "refreshToken":"${DART_REFRESH_TOKEN}",
  "tokenEndpoint": "https://accounts.google.com/o/oauth2/token",
  "scopes": [
    "openid",
    "https://www.googleapis.com/auth/userinfo.email"
  ],
  "expiration": 1611594593613
}
EOF
  set +x

  (cd "${dir}"; VERSION=${RAW_VERSION} command dart pub publish --force)
}

rust() {
  dir="clients/${PROJECT}/rust"

  cargo login "${CARGO_TOKEN}"

  (cd "${dir}"; VERSION=${RAW_VERSION} command cargo publish --allow-dirty)
}

for i in {1..10}; do python && break || sleep 15; done
for i in {1..10}; do ruby && break || sleep 15; done
for i in {1..10}; do golang && break || sleep 15; done
for i in {1..10}; do php && break || sleep 15; done
for i in {1..10}; do typescript && break || sleep 15; done
for i in {1..10}; do dart && break || sleep 15; done
for i in {1..10}; do rust && break || sleep 15; done
for i in {1..10}; do java && break || sleep 15; done
for i in {1..10}; do dotnet && break || sleep 15; done

upstream || true
