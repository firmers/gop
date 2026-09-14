#
#   Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
#   Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
#
#

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
