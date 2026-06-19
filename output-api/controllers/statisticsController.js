const {calculateStatistics} = require('../services/statisticsService');

function getStatistics(req, res) {
    const { q, r } = req.body;

    if (!Array.isArray(q) || !Array.isArray(r)) {
        return res.status(400).json({ error: 'q and r matrices are required' });
    }

    const result = calculateStatistics(q, r);

    return res.json(result);
}

module.exports = {
    getStatistics,
};