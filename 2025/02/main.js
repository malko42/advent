import { readFile } from 'fs/promises';

async function load_data() {
  return await readFile('input.txt', 'utf-8');
  // return await readFile('sample.txt', 'utf-8');
}

function parse_data(data) {
  // split by newline, discard empty lines
  const lines = data.split('\n').map(line => line.trim()).filter(line => line.length > 0);
  console.log(`Loaded ${lines.length} lines of data.`);
  return lines.map(line => {
    let line_array = Array.from(line);
    return line;
  });
}

function process_1(lines) {
}

function process_2(lines) {
}

async function main() {
  const data = await load_data();
  const lines = parse_data(data);
  const result_1 = process_1(lines);
  console.log(`Result 1: ${result_1}`);
  const result_2 = process_2(lines);
  console.log(`Result 2: ${result_2}`);
}

await main();
