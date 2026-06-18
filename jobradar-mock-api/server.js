const express = require('express');
const cors = require('cors');
const path = require('path');

// Memuat data JSON hasil scraping Data Pipeline
const jobsData = require('./jobradar_mock_db.json'); 

const app = express();

// Middleware
app.use(cors()); // Membuka akses untuk frontend React
app.use(express.json());

// Endpoint Utama
app.get('/api/jobs', (req, res) => {
    console.log(`[GET] Permintaan masuk ke /api/jobs dari ${req.ip}`);
    
    // Opsional: Simulasi delay jaringan (misal: 800ms) agar terasa nyata
    // Membantu frontend menguji loading state spinner/skeleton
    setTimeout(() => {
        res.status(200).json({
            success: true,
            total_data: jobsData.length,
            data: jobsData
        });
    }, 800);
});

// Endpoint untuk mendapatkan job spesifik (opsional)
app.get('/api/jobs/:id', (req, res) => {
    // Karena saat ini ID belum distandarisasi di normalizer, 
    // kita gunakan index array sementara
    const jobId = parseInt(req.params.id);
    if (jobId >= 0 && jobId < jobsData.length) {
        res.status(200).json({
            success: true,
            data: jobsData[jobId]
        });
    } else {
        res.status(404).json({ success: false, message: "Pekerjaan tidak ditemukan" });
    }
});

const PORT = process.env.PORT || 5000;

app.listen(PORT, () => {
    console.log(`🚀 Mock API Server berhasil berjalan!`);
    console.log(`📡 Endpoint utama: http://localhost:${PORT}/api/jobs`);
});