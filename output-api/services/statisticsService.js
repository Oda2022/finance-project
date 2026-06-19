
function flattenMatrices(matrices) {
  return matrices.flat(2);
}

function isDiagonal(matrix) {
    return matrix.every((row, i) => row.every((value, j) => i === j || Math.abs(value) < 1e-10));
}

function calculateStatistics(q, r) {
    const matrices = [q, r];
    const flattened = flattenMatrices(matrices);

    const sum = flattened.reduce((acc, value) => acc + value, 0);

    return {
        max: Math.max(...flattened),
        min: Math.min(...flattened),
        average: sum / flattened.length,
        sum,
        hasDiagonalMatrix: matrices.some(isDiagonal),
    }
}

module.exports = {
    calculateStatistics,
};