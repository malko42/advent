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
    const rot = line_array.shift();
    const deg = parseInt(line_array.join(''));
    return {
      rot,
      deg
    }
  });
}

function process_1(lines) {
  let zero = 0;
  let start_pos = 50;

  for (const line of lines) {
    if (line.rot === 'L') {
      start_pos -= line.deg;
    }

    if (line.rot === 'R') {
      start_pos += line.deg;
    }

    start_pos = start_pos % 100;

    if (start_pos < 0) {
      start_pos += 100;
    }

    if (start_pos == 0) {
      zero += 1;
    }
  }

  return zero;
}

function process_2(lines) {
  let zero = 0;
  let start_pos = 50;

  for (const line of lines) {
    if (line.rot === 'L') {
      start_pos -= line.deg;
    }

    if (line.rot === 'R') {
      start_pos += line.deg;
    }

    zero += Math.floor(Math.abs(start_pos) / 100);

    start_pos = start_pos % 100;

    if (start_pos < 0) {
      zero += 1;
      start_pos += 100;
    }
  }

  return zero;
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
