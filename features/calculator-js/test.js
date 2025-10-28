// Simple test runner without external dependencies
const { add, subtract, multiply, divide, power, modulo } = require('./calculator.js');

let passed = 0;
let failed = 0;

function assert(condition, message) {
  if (condition) {
    console.log(`✓ ${message}`);
    passed++;
  } else {
    console.error(`✗ ${message}`);
    failed++;
  }
}

function assertEqual(actual, expected, message) {
  if (actual === expected) {
    console.log(`✓ ${message}`);
    passed++;
  } else {
    console.error(`✗ ${message} - Expected ${expected}, got ${actual}`);
    failed++;
  }
}

console.log('Running unit tests...\n');

// Test add function
assertEqual(add(2, 3), 5, 'add(2, 3) should return 5');
assertEqual(add(-1, 1), 0, 'add(-1, 1) should return 0');
assertEqual(add(0, 0), 0, 'add(0, 0) should return 0');

// Test subtract function
assertEqual(subtract(5, 3), 2, 'subtract(5, 3) should return 2');
assertEqual(subtract(0, 5), -5, 'subtract(0, 5) should return -5');

// Test multiply function
assertEqual(multiply(3, 4), 12, 'multiply(3, 4) should return 12');
assertEqual(multiply(-2, 3), -6, 'multiply(-2, 3) should return -6');

// Test divide function
assertEqual(divide(10, 2), 5, 'divide(10, 2) should return 5');
assertEqual(divide(7, 2), 3.5, 'divide(7, 2) should return 3.5');

// Test divide by zero
try {
  divide(10, 0);
  console.error('✗ divide by zero should throw an error');
  failed++;
} catch (e) {
  console.log('✓ divide by zero correctly throws an error');
  passed++;
}

// Test power function
assertEqual(power(2, 3), 8, 'power(2, 3) should return 8');
assertEqual(power(5, 2), 25, 'power(5, 2) should return 25');
assertEqual(power(10, 0), 1, 'power(10, 0) should return 1');

// Test modulo function
assertEqual(modulo(10, 3), 1, 'modulo(10, 3) should return 1');
assertEqual(modulo(15, 4), 3, 'modulo(15, 4) should return 3');
assertEqual(modulo(8, 2), 0, 'modulo(8, 2) should return 0');

console.log(`\n${'='.repeat(40)}`);
console.log(`Total: ${passed + failed} tests`);
console.log(`Passed: ${passed}`);
console.log(`Failed: ${failed}`);
console.log(`${'='.repeat(40)}`);

process.exit(failed > 0 ? 1 : 0);
