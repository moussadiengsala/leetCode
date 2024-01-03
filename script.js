var fizzBuzz = function (n) {
    let arr = [];
    for (let i = 1; i <= n; i++) {
        if (i % 3 === 0 && i % 5 === 0) {
            arr.push("FizzBuzz");
        } else if (i % 3 === 0) {
            arr.push("Fizz");
        } else if (i % 5 === 0) {
            arr.push("Buzz");
        } else {
            arr.push(`${i}`);
        }
    }
    return arr;
};

// Constraints:

//     1 <= n <= 104

// console.log(fizzBuzz(5));

/**
 * @param {number} left
 * @param {number} right
 * @return {number[]}
 */
var selfDividingNumbers = function (left, right) {
    arr = [];
    for (var i = left; i <= right; i++) {
        let a = `${i}`;
        let digits = a.split("");

        if (digits.every((digit) => i % digit === 0)) arr.push(i);
    }
    return arr;
};

// console.log(selfDividingNumbers(12, 14))

var numIdenticalPairs = function (nums) {
    let pairs = 0;
    for (var i = 0; i < nums.length; i++) {
        for (var j = 0; j < nums.length; j++) {
            if (i < j && nums[i] === nums[j]) pairs++;
        }
    }
    return pairs;
};

/**
 * @param {number} celsius
 * @return {number[]}
 */
var convertTemperature = function (celsius) {
    // Kelvin = Celsius + 273.15
    // Fahrenheit = Celsius * 1.80 + 32.00

    let kelvin = (celsius + 273.15).toFixed(5);
    let fahrenheit = (celsius * 1.8 + 32.0).toFixed(5);

    return [kelvin, fahrenheit];
};

// console.log(convertTemperature(36.50)) // should be that: [309.65000,97.70000]

var shuffle = function (nums, n) {
    let shufflearr = [];
    for (let i = 0; i < n; i++) {
        shufflearr.push(nums[i], nums[n + i]);
    }

    return shufflearr;
};

// console.log(shuffle([2,5,1,3,4,7], 3))

var numberOfEmployeesWhoMetTarget = function (hours, target) {
    return hours.filter((hour) => hour >= target).length;
};

/**
 * @param {number} n
 * @return {number}
 */
var smallestEvenMultiple = function (n) {
    for (var i = 1; ; i++) {
        if (i % 2 === 0 && i % n === 0) return i;
    }
};

/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function (x) {
    let y = +`${x}`.split("").reverse().join("");
    return !Number.isNaN(y) && y == x;
};

// console.log(isPalindrome(121))

/**
 * @param {string} word1
 * @param {string} word2
 * @return {string}
 */
var mergeAlternately = function (word1, word2) {
    let limit = word1.length > word2.length ? word1.length : word2.length;
    let letters = [];
    for (let i = 0; i < limit; i++) {
        letters.push(word1[i], word2[i]);
    }
    return letters.filter((letter) => letter !== undefined).join("");
};

// console.log(mergeAlternately("ab", "pqrs"))

/**
 * @param {number} n
 * @return {number}
 */
var countPrimes = function (n) {
    if (n <= 2) return 0;

    let isPrime = new Array(n).fill(true);
    console.log(isPrime);
    isPrime[0] = false;
    isPrime[1] = false;

    for (let i = 2; i * i < n; i++) {
        if (isPrime[i]) {
            for (let j = i * i; j < n; j += i) {
                isPrime[j] = false;
            }
        }
    }

    return isPrime.filter((prime) => prime).length;
};

console.log(countPrimes(10));
