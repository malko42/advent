import { readFile } from 'fs/promises';

async function load_data() {
  return await readFile('input.txt', 'utf-8');
  // return await readFile('sample.txt', 'utf-8');
}

function parse_data(data) {
  // split by comma
  const ranges = data.split(',').map(range => range.trim()).map(range => {
    const tmp = range.split('-');
    const result = { start: tmp[0], end: tmp[1] };
    if (isNaN(parseInt(result.start)) || isNaN(parseInt(result.end))) {
      throw new Error(`Invalid range: ${range}`);
    }
    if (parseInt(result.start) > parseInt(result.end)) {
      throw new Error(`Start greater than end: ${range}`);
    }
    return result;
  });
  console.log(`Loaded ${ranges.length} ranges.`);
  return ranges;
}

function is_invalid(number) {
  // number is a string
  if (number.length % 2 !== 0) {
    return false;
  }

  const left = number.slice(0, number.length / 2);
  const right = number.slice(number.length / 2);

  if (parseInt(left) == parseInt(right)) {
    return true;
  }

  return false;
}

function is_repeating(number) {
  if (number.length < 2) {
    return false;
  }
  const half = Math.floor(number.length / 2);
  const left = number.slice(0, half);
  const right = number.slice(half);
  if (left == right) {
    return true;
  }

  for (let i = half; i > 0; i--) {
    const reg = new RegExp(`.{1,${i}}`, 'g');
    const chunks = number.match(reg);
    const test = chunks.every((val, _, arr) => val === arr[0]);
    if (test) {
      return true;
    }
  }
  return false;
}
function process_1(ranges) {

  let sum = 0;
  for (const range of ranges) {
    for (let i = parseInt(range.start); i <= parseInt(range.end); i++) {
      if (is_invalid(i.toString())) {
        sum += i;
      }
    }
  }
  return sum;
}

function process_2(ranges) {
  let sum = 0;
  for (const range of ranges) {
    console.log(`Processing range: ${range.start}-${range.end}`);
    for (let i = parseInt(range.start); i <= parseInt(range.end); i++) {
      if (is_repeating(i.toString())) {
        console.log(`Found repeating number: ${i}`);
        sum += i;
      }
    }
  }
  return sum;
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
