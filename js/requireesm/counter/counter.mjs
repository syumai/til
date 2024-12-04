// ES modules側 (counter.mjs)
let count = 0;
export const currentCount = () => count;
export const increment = () => ++count;
export const decrement = () => --count;
