#
#   Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       https://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

.PHONY: i
i:
	flutter pub get

.PHONY: p
p:
	flutter clean
	flutter build apk

.PHONY: b
b:
	dart run build_runner build --force-jit

.PHONY: c
c:
	flutter build ios --no-codesign --config-only
	flutter build apk --no-codesign --config-only

.PHONY: g
g:
	dart run build_runner build

.PHONY: u
u:
	dart pub publish --skip-validation

.PHONY: l10n
l10n:
	flutter gen-l10n
