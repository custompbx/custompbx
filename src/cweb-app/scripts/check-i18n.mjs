import {readFile, readdir} from 'node:fs/promises';
import {resolve} from 'node:path';

const root = resolve('src/assets/i18n');
const expectedLocales = ['en', 'fr', 'de', 'es', 'pt-BR', 'it', 'tr', 'ru', 'ar', 'fa', 'hi', 'zh-Hans', 'ja', 'ko'];

function keys(value, prefix = '') {
  return Object.entries(value).flatMap(([key, child]) => {
    const path = prefix ? `${prefix}.${key}` : key;
    return child && typeof child === 'object' && !Array.isArray(child) ? keys(child, path) : [path];
  });
}

const files = (await readdir(root)).filter(file => file.endsWith('.json')).sort();
const expectedFiles = expectedLocales.map(locale => `${locale}.json`).sort();
if (JSON.stringify(files) !== JSON.stringify(expectedFiles)) {
  throw new Error(`Locale files differ. Expected ${expectedFiles.join(', ')}, got ${files.join(', ')}`);
}

const catalogs = new Map();
for (const locale of expectedLocales) {
  const file = resolve(root, `${locale}.json`);
  catalogs.set(locale, JSON.parse(await readFile(file, 'utf8')));
}

const englishKeys = keys(catalogs.get('en')).sort();
for (const [locale, catalog] of catalogs) {
  const localeKeys = keys(catalog).sort();
  const missing = englishKeys.filter(key => !localeKeys.includes(key));
  const extra = localeKeys.filter(key => !englishKeys.includes(key));
  if (missing.length || extra.length) {
    throw new Error(`${locale}: missing [${missing.join(', ')}], extra [${extra.join(', ')}]`);
  }
}

console.log(`Validated ${expectedLocales.length} locale catalogs with ${englishKeys.length} keys.`);
