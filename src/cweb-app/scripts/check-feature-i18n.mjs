import fs from 'node:fs';
import path from 'node:path';
import ts from 'typescript';

const sourcePath = path.resolve('src/app/i18n/feature-translations.ts');
const source = fs.readFileSync(sourcePath, 'utf8');
const compiled = ts.transpileModule(source, {
  compilerOptions: {module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2022},
}).outputText;
const module = {exports: {}};
new Function('exports', 'module', 'require', compiled)(module.exports, module, () => ({}));

const {
  FEATURE_TRANSLATIONS_ENGLISH: english,
  FEATURE_TRANSLATION_OVERRIDES: overrides,
  FEATURE_TRANSLATION_COMPLETIONS: completions,
  FEATURE_TRANSLATION_UI_COMPLETIONS: uiCompletions,
} = module.exports;
const localeCodes = ['fr', 'de', 'es', 'pt-BR', 'it', 'tr', 'ru', 'ar', 'fa', 'hi', 'zh-Hans', 'ja', 'ko'];

function leafPaths(value, prefix = '') {
  const paths = [];
  for (const [key, child] of Object.entries(value ?? {})) {
    const current = prefix ? `${prefix}.${key}` : key;
    if (child && typeof child === 'object' && !Array.isArray(child)) {
      paths.push(...leafPaths(child, current));
    } else {
      paths.push(current);
    }
  }
  return paths;
}

function merge(base, extra) {
  const result = {...base};
  for (const [key, value] of Object.entries(extra ?? {})) {
    result[key] = value && typeof value === 'object' && !Array.isArray(value)
      ? merge(result[key] ?? {}, value)
      : value;
  }
  return result;
}

const expected = new Set(leafPaths(english));
const failures = [];
for (const locale of localeCodes) {
  const native = merge(merge(overrides[locale] ?? {}, completions[locale] ?? {}), uiCompletions[locale] ?? {});
  const present = new Set(leafPaths(native));
  const missing = [...expected].filter(key => !present.has(key));
  if (missing.length) {
    failures.push(`${locale}: ${missing.join(', ')}`);
  }
}

if (failures.length) {
  console.error(`Feature translation coverage failed:\n${failures.join('\n')}`);
  process.exit(1);
}

console.log(`${localeCodes.length} feature catalogs cover all ${expected.size} feature labels.`);
