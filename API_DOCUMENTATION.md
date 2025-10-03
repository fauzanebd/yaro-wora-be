# Yaro Wora Tourism Website - Backend API Documentation

This document outlines the required backend endpoints to support the Yaro Wora tourism website frontend.

## Base URL

```
https://api.yarowora.com/v1
```

## Authentication

- For public endpoints: No authentication required
- For admin endpoints: Bearer token authentication (see ADMIN_API_DOCUMENTATION.md)
- Content-Type: `application/json`

---

## 🏠 Main Page

### GET /carousel

**Description**: Get carousel slides for the hero section

```json
{
  "data": [
    {
      "id": 1,
      "title": "Let's Discover Yaro Wora: An Authentic Culture and Nature",
      "subtitle": "Where untold stories and hidden beauty await to be discovered",
      "image_url": "https://api.yarowora.com/images/carousel/slide1.jpg",
      "alt_text": "Beautiful mountain landscape of Yaro Wora",
      "order": 1,
      "is_active": true
    }
  ],
  "meta": {
    "total": 3,
    "auto_play_interval": 6000
  }
}
```

### GET /attractions

**Description**: Get featured attractions for the main page

```json
{
  "data": [
    {
      "id": "traditional-village",
      "title": "Traditional Village Tour",
      "short_description": "Explore authentic Sumba architecture and daily village life.",
      "full_description": "Immerse yourself in the living heritage...",
      "image_url": "https://api.yarowora.com/images/attractions/village.jpg",
      "highlights": [
        "Traditional Architecture",
        "Local Crafts",
        "Cultural Stories"
      ],
      "duration": "3-4 hours",
      "difficulty": "easy",
      "price_range": "50000-100000",
      "is_featured": true
    }
  ],
  "meta": {
    "total": 6,
    "featured_count": 3
  }
}
```

### GET /pricing

**Description**: Get entrance fee pricing for different visitor types

```json
{
  "data": {
    "domestic": {
      "title": "Domestic",
      "subtitle": "Entrance Fee",
      "adult_price": 60000,
      "infant_price": 30000,
      "currency": "IDR",
      "description": "For Indonesian citizens"
    },
    "locals_sumba": {
      "title": "Locals Sumba",
      "subtitle": "Entrance Fee",
      "adult_price": 40000,
      "infant_price": 20000,
      "currency": "IDR",
      "description": "For local Sumba residents"
    },
    "foreigner": {
      "title": "Foreigner",
      "subtitle": "Entrance Fee",
      "adult_price": 100000,
      "infant_price": 50000,
      "currency": "IDR",
      "description": "For international visitors"
    }
  },
  "last_updated": "2024-01-15T10:30:00Z"
}
```

---

## 👤 Profile Page

### GET /profile

**Description**: Get village profile information including vision, mission, and objectives

**Query Parameters**:

- `lang`: `en|id` (Language for localized content)

**Headers**:

- `Accept-Language`: `en|id` (Alternative way to specify language)

```json
{
  "data": {
    "title": "Desa Patiala Bawa - Kampung Yaro Wora",
    "description": "Yaro Wora is a traditional village nestled in the heart of Sumba, Indonesia, where ancient customs and pristine nature converge to create an authentic cultural experience.\n\nOur village represents a living heritage where time-honored traditions continue to thrive amidst breathtaking natural landscapes. Here, visitors can witness the authentic Sumba way of life, participate in traditional ceremonies, and explore unspoiled natural beauty.\n\nThe community of Yaro Wora is committed to preserving our cultural identity while welcoming respectful visitors who seek to understand and appreciate our unique way of life.",
    "vision": {
      "title": "Visi",
      "content": "**Menjadi desa wisata budaya terdepan di Sumba yang memadukan kearifan lokal dengan pembangunan berkelanjutan.**\n\nKami bercita-cita untuk menjadi destinasi wisata yang tidak hanya menawarkan keindahan alam dan kekayaan budaya, tetapi juga menjadi contoh bagaimana masyarakat lokal dapat mempertahankan identitas budaya mereka sambil berkembang secara ekonomi.\n\nVisi kami mencakup:\n- Pelestarian tradisi dan budaya Sumba yang autentik\n- Pemberdayaan masyarakat lokal melalui pariwisata berkelanjutan\n- Perlindungan lingkungan dan ekosistem alami\n- Peningkatan kesejahteraan masyarakat desa"
    },
    "mission": {
      "title": "Misi",
      "content": "**Mengembangkan potensi wisata desa dengan tetap menjaga kelestarian budaya dan lingkungan.**\n\n### Misi Utama Kami:\n\n1. **Pelestarian Budaya**\n   - Mempertahankan tradisi dan adat istiadat Sumba\n   - Mengajarkan generasi muda tentang warisan budaya\n   - Mendokumentasikan dan melestarikan cerita-cerita tradisional\n\n2. **Pengembangan Pariwisata Berkelanjutan**\n   - Menyediakan pengalaman wisata yang autentik dan berkesan\n   - Mengembangkan fasilitas wisata yang ramah lingkungan\n   - Melatih masyarakat lokal sebagai pemandu wisata\n\n3. **Pemberdayaan Ekonomi**\n   - Menciptakan lapangan kerja bagi masyarakat lokal\n   - Mengembangkan usaha kecil dan menengah di bidang pariwisata\n   - Meningkatkan pendapatan masyarakat melalui ekonomi kreatif\n\n4. **Konservasi Lingkungan**\n   - Menjaga kelestarian alam dan biodiversitas\n   - Menerapkan praktik ramah lingkungan\n   - Eduaksi wisatawan tentang pentingnya konservasi"
    },
    "objectives": {
      "title": "Tujuan",
      "content": "**Menciptakan ekosistem pariwisata yang berkelanjutan dan bermanfaat bagi semua pihak.**\n\n### Tujuan Jangka Pendek (1-2 Tahun):\n- Meningkatkan jumlah kunjungan wisatawan sebesar 50%\n- Mengembangkan 5 paket wisata baru yang menampilkan berbagai aspek budaya Sumba\n- Melattih 20 pemandu wisata lokal yang bersertifikat\n- Membangun pusat informasi wisata yang representatif\n\n### Tujuan Jangka Menengah (3-5 Tahun):\n- Menjadikan Yaro Wora sebagai destinasi wisata budaya unggulan di Sumba\n- Mengembangkan homestay dan fasilitas akomodasi berkelanjutan\n- Menciptakan 100 lapangan kerja langsung di sektor pariwisata\n- Meluncurkan program pertukaran budaya internasional\n\n### Tujuan Jangka Panjang (5+ Tahun):\n- Mencapai status sebagai desa wisata budaya berstandar internasional\n- Mengembangkan pusat penelitian budaya Sumba\n- Menjadi model pengembangan pariwisata berkelanjutan untuk desa-desa lain\n- Melestarikan dan mengembangkan warisan budaya untuk generasi mendatang\n\n**Dengan komitmen bersama, kami yakin dapat mencapai semua tujuan ini sambil tetap mempertahankan keaslian dan keunikan budaya Sumba.**"
    },
    "featured_images": [
      {
        "url": "https://api.yarowora.com/images/profile/village-overview.jpg",
        "caption": "Overview of Yaro Wora village",
        "alt_text": "Traditional houses with mountain backdrop"
      },
      {
        "url": "https://api.yarowora.com/images/profile/cultural-ceremony.jpg",
        "caption": "Traditional ceremony in progress",
        "alt_text": "Villagers performing traditional dance"
      }
    ],
    "last_updated": "2024-01-15T08:00:00Z"
  }
}
```

---

## 🗺️ Destinations Page

### GET /destinations

**Description**: Get all destinations with filtering options

**Query Parameters**:

- `type`: `main|other` (filter by destination type - main is featured destination, other are regular cards)
- `category`: `nature|culture|heritage|agriculture|adventure`
- `limit`: number of results per page
- `offset`: pagination offset
- `lang`: `en|id` (language for localized content)

```json
{
  "data": [
    {
      "id": "traditional-village-center",
      "title": "Traditional Village Center",
      "title_id": "Pusat Desa Tradisional",
      "description": "The heart of Yaro Wora, featuring authentic Sumba architecture, traditional houses with iconic pointed roofs, and the main community gathering space where cultural ceremonies take place.",
      "description_id": "Jantung Yaro Wora, menampilkan arsitektur asli Sumba, rumah tradisional beratap runcing ikonik, dan ruang berkumpul komunitas utama tempat upacara budaya berlangsung.",
      "image_url": "https://api.yarowora.com/images/destinations/village-center.jpg",
      "thumbnail_url": "https://api.yarowora.com/images/destinations/thumbs/village-center.jpg",
      "category": "culture",
      "category_id": "budaya",
      "type": "main",
      "location": {
        "latitude": -9.6234,
        "longitude": 119.3456,
        "address": "Yaro Wora Village Center, East Sumba"
      },
      "highlights": [
        "Traditional Sumba Architecture",
        "Cultural Ceremony Grounds",
        "Community Gathering Space",
        "Historic Significance"
      ],
      "highlights_id": [
        "Arsitektur Tradisional Sumba",
        "Tempat Upacara Budaya",
        "Ruang Berkumpul Komunitas",
        "Signifikansi Bersejarah"
      ],
      "best_visit_time": "morning",
      "duration": "2-4 hours",
      "difficulty": "easy",
      "is_featured": true,
      "sort_order": 1
    }
  ],
  "meta": {
    "total": 7,
    "main_destinations": 1,
    "other_destinations": 6,
    "categories": [
      {
        "id": "culture",
        "name": "Cultural Heritage",
        "name_id": "Warisan Budaya",
        "count": 3
      },
      {
        "id": "nature",
        "name": "Nature",
        "name_id": "Alam",
        "count": 3
      },
      {
        "id": "agriculture",
        "name": "Agriculture",
        "name_id": "Pertanian",
        "count": 1
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 10,
      "total_pages": 1
    }
  }
}
```

### GET /destinations/main

**Description**: Get the main featured destination specifically

**Query Parameters**:

- `lang`: `en|id` (language for localized content)

```json
{
  "data": {
    "id": "traditional-village-center",
    "title": "Traditional Village Center",
    "title_id": "Pusat Desa Tradisional",
    "description": "The heart of Yaro Wora, featuring authentic Sumba architecture...",
    "description_id": "Jantung Yaro Wora, menampilkan arsitektur asli Sumba...",
    "image_url": "https://api.yarowora.com/images/destinations/village-center.jpg",
    "hero_image_url": "https://api.yarowora.com/images/destinations/village-center-hero.jpg",
    "category": "culture",
    "category_id": "budaya",
    "highlights": [
      "Traditional Sumba Architecture",
      "Cultural Ceremony Grounds",
      "Community Gathering Space",
      "Historic Significance"
    ],
    "highlights_id": [
      "Arsitektur Tradisional Sumba",
      "Tempat Upacara Budaya",
      "Ruang Berkumpul Komunitas",
      "Signifikansi Bersejarah"
    ],
    "type": "main",
    "is_featured": true
  }
}
```

### GET /destinations/{id}

**Description**: Get detailed information about a specific destination for the modal

```json
{
  "data": {
    "id": "traditional-village-center",
    "title": "Traditional Village Center",
    "title_id": "Pusat Desa Tradisional",
    "description": "Detailed description of the destination with comprehensive information about what visitors can expect...",
    "description_id": "Deskripsi detail destinasi dengan informasi komprehensif tentang apa yang dapat diharapkan pengunjung...",
    "full_content": "# Traditional Village Center\n\nThe Traditional Village Center serves as the cultural heart of Yaro Wora...",
    "full_content_id": "# Pusat Desa Tradisional\n\nPusat Desa Tradisional berfungsi sebagai jantung budaya Yaro Wora...",
    "images": [
      {
        "url": "https://api.yarowora.com/images/destinations/village-center-1.jpg",
        "caption": "Main village houses with traditional architecture",
        "caption_id": "Rumah-rumah desa utama dengan arsitektur tradisional",
        "is_primary": true
      },
      {
        "url": "https://api.yarowora.com/images/destinations/village-center-2.jpg",
        "caption": "Community gathering during cultural ceremony",
        "caption_id": "Pertemuan komunitas saat upacara budaya",
        "is_primary": false
      }
    ],
    "location": {
      "latitude": -9.6234,
      "longitude": 119.3456,
      "address": "Yaro Wora Village Center, East Sumba",
      "address_id": "Pusat Desa Yaro Wora, Sumba Timur",
      "google_maps_url": "https://maps.google.com/?q=-9.6234,119.3456"
    },
    "highlights": [
      "Traditional Sumba Architecture",
      "Cultural Ceremony Grounds",
      "Community Gathering Space",
      "Historic Significance"
    ],
    "highlights_id": [
      "Arsitektur Tradisional Sumba",
      "Tempat Upacara Budaya",
      "Ruang Berkumpul Komunitas",
      "Signifikansi Bersejarah"
    ],
    "facilities": [
      "Guided Tours",
      "Traditional Performances",
      "Restrooms",
      "Photo Opportunities"
    ],
    "facilities_id": [
      "Tur Berpemandu",
      "Pertunjukan Tradisional",
      "Toilet",
      "Peluang Foto"
    ],
    "best_visit_time": {
      "period": "morning",
      "description": "Best visited in the morning when the light is perfect for photography and traditional activities are most active",
      "description_id": "Paling baik dikunjungi di pagi hari ketika cahaya sempurna untuk fotografi dan aktivitas tradisional paling aktif"
    },
    "duration": {
      "recommended": "2-4 hours",
      "description": "Complete exploration with cultural activities and photography",
      "description_id": "Eksplorasi lengkap dengan aktivitas budaya dan fotografi"
    },
    "difficulty": "easy",
    "accessibility": "wheelchair_friendly",
    "entrance_fee": null,
    "opening_hours": {
      "monday": "08:00-17:00",
      "tuesday": "08:00-17:00",
      "wednesday": "08:00-17:00",
      "thursday": "08:00-17:00",
      "friday": "08:00-17:00",
      "saturday": "08:00-16:00",
      "sunday": "09:00-15:00"
    },
    "booking_required": false,
    "contact_info": {
      "phone": "+62 903 009 909",
      "whatsapp": "+62 903 009 909",
      "email": "visit@yarowora.com"
    }
  }
}
```

---

## 🖼️ Gallery Page

### GET /gallery

**Description**: Get gallery images with Pinterest-style layout support

**Query Parameters**:

- `category`: `nature|culture|village|activities`
- `limit`: number of images per page
- `offset`: pagination offset
- `search`: search query for title, description, or tags

```json
{
  "data": [
    {
      "id": "img_001",
      "title": "Sunrise over Yaro Wora",
      "description": "Golden sunrise illuminating the traditional village landscape with morning mist",
      "image_url": "https://api.yarowora.com/images/gallery/sunrise.jpg",
      "thumbnail_url": "https://api.yarowora.com/images/gallery/thumbs/sunrise.jpg",
      "category": "nature",
      "dimensions": {
        "width": 1200,
        "height": 800,
        "aspect_ratio": 1.5
      },
      "photographer": "Local Guide Team",
      "date_created": "2024-01-05T06:30:00Z",
      "location": "Yaro Wora Viewpoint",
      "tags": ["sunrise", "landscape", "village", "mist"]
    }
  ],
  "meta": {
    "total": 150,
    "categories": ["nature", "culture", "village", "activities"],
    "pagination": {
      "current_page": 1,
      "per_page": 20,
      "total_pages": 8
    }
  }
}
```

### GET /gallery/categories

**Description**: Get available gallery categories with image counts

```json
{
  "data": [
    {
      "id": "nature",
      "name": "Nature",
      "name_id": "Alam",
      "description": "Natural landscapes and wildlife photography",
      "description_id": "Fotografi lanskap alam dan satwa liar",
      "count": 45,
      "color": "#22c55e",
      "icon": "nature"
    },
    {
      "id": "culture",
      "name": "Culture",
      "name_id": "Budaya",
      "description": "Cultural ceremonies and traditional activities",
      "description_id": "Upacara budaya dan aktivitas tradisional",
      "count": 38,
      "color": "#8a0604",
      "icon": "culture"
    },
    {
      "id": "village",
      "name": "Village Life",
      "name_id": "Kehidupan Desa",
      "description": "Daily life and village architecture",
      "description_id": "Kehidupan sehari-hari dan arsitektur desa",
      "count": 42,
      "color": "#dc2626",
      "icon": "village"
    },
    {
      "id": "activities",
      "name": "Activities",
      "name_id": "Aktivitas",
      "description": "Tourism activities and experiences",
      "description_id": "Aktivitas pariwisata dan pengalaman",
      "count": 25,
      "color": "#586d12",
      "icon": "activities"
    }
  ],
  "meta": {
    "total_categories": 4,
    "total_images": 150
  }
}
```

### GET /gallery/{id}

**Description**: Get detailed information about a specific gallery image

```json
{
  "data": {
    "id": "img_001",
    "title": "Sunrise over Yaro Wora",
    "description": "Full detailed description of the image and its significance...",
    "image_url": "https://api.yarowora.com/images/gallery/sunrise.jpg",
    "high_res_url": "https://api.yarowora.com/images/gallery/high-res/sunrise.jpg",
    "metadata": {
      "camera": "Canon EOS R5",
      "settings": "f/8, 1/125s, ISO 100",
      "location": "Yaro Wora Viewpoint"
    },
    "related_images": ["img_002", "img_003"]
  }
}
```

---

## 📋 Regulations Page

### GET /regulations

**Description**: Get regulations and guidelines for visiting Yaro Wora village

**Query Parameters**:

- `category` (optional): Filter by category (`general`, `cultural`, `operational`, `safety`, `environmental`, `accommodation`, `accessibility`)
- `search` (optional): Search within regulation content
- `is_active` (optional): Filter active regulations (default: `true`)
- `page` (optional): Page number for pagination (default: 1)
- `limit` (optional): Items per page (default: 20, max: 100)

```json
{
  "data": [
    {
      "id": "general-rules",
      "category": "general",
      "question": "What are the general rules for visiting Yaro Wora village?",
      "answer": "Visitors are expected to respect local customs and traditions. Please dress modestly, especially when visiting sacred areas. Photography is allowed in most areas, but please ask permission before taking photos of people. Smoking and alcohol consumption are prohibited in certain areas. Please follow your guide's instructions at all times.",
      "priority": 1,
      "is_active": true,
      "created_at": "2024-01-10T08:00:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    },
    {
      "id": "entry-requirements",
      "category": "general",
      "question": "What are the entry requirements and entrance fees?",
      "answer": "All visitors must register at the village entrance and pay the appropriate entrance fee based on their category (domestic, local Sumba, or foreigner). Valid identification is required. Groups of 10 or more must book in advance. Children under 5 enter free with adult supervision.",
      "priority": 2,
      "is_active": true,
      "created_at": "2024-01-10T08:00:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  ],
  "meta": {
    "total": 15,
    "page": 1,
    "limit": 20,
    "total_pages": 1,
    "categories": [
      {
        "key": "general",
        "name": "General Rules",
        "count": 3
      },
      {
        "key": "cultural",
        "name": "Cultural Guidelines",
        "count": 4
      },
      {
        "key": "operational",
        "name": "Operations",
        "count": 2
      },
      {
        "key": "safety",
        "name": "Safety & Emergency",
        "count": 2
      },
      {
        "key": "environmental",
        "name": "Environmental",
        "count": 2
      },
      {
        "key": "accommodation",
        "name": "Accommodation",
        "count": 1
      },
      {
        "key": "accessibility",
        "name": "Accessibility",
        "count": 1
      }
    ]
  }
}
```

### GET /regulations/categories

**Description**: Get available regulation categories with statistics

```json
{
  "data": [
    {
      "key": "general",
      "name": "General Rules",
      "name_id": "Aturan Umum",
      "description": "Basic rules and requirements for all visitors",
      "description_id": "Aturan dasar dan persyaratan untuk semua pengunjung",
      "count": 3,
      "icon": "rules",
      "color": "#6b7280"
    },
    {
      "key": "cultural",
      "name": "Cultural Guidelines",
      "name_id": "Panduan Budaya",
      "description": "Important cultural etiquette and customs",
      "description_id": "Etiket budaya penting dan adat istiadat",
      "count": 4,
      "icon": "culture",
      "color": "#8a0604"
    },
    {
      "key": "operational",
      "name": "Operations",
      "name_id": "Operasional",
      "description": "Operating hours, booking, and administrative policies",
      "description_id": "Jam operasional, pemesanan, dan kebijakan administratif",
      "count": 2,
      "icon": "clock",
      "color": "#059669"
    },
    {
      "key": "safety",
      "name": "Safety & Emergency",
      "name_id": "Keselamatan & Darurat",
      "description": "Safety protocols and emergency procedures",
      "description_id": "Protokol keselamatan dan prosedur darurat",
      "count": 2,
      "icon": "shield",
      "color": "#dc2626"
    },
    {
      "key": "environmental",
      "name": "Environmental",
      "name_id": "Lingkungan",
      "description": "Environmental conservation and protection guidelines",
      "description_id": "Pedoman konservasi dan perlindungan lingkungan",
      "count": 2,
      "icon": "leaf",
      "color": "#16a34a"
    },
    {
      "key": "accommodation",
      "name": "Accommodation",
      "name_id": "Akomodasi",
      "description": "Homestay and accommodation guidelines",
      "description_id": "Panduan homestay dan akomodasi",
      "count": 1,
      "icon": "home",
      "color": "#7c3aed"
    },
    {
      "key": "accessibility",
      "name": "Accessibility",
      "name_id": "Aksesibilitas",
      "description": "Accessibility information for visitors with disabilities",
      "description_id": "Informasi aksesibilitas untuk pengunjung dengan disabilitas",
      "count": 1,
      "icon": "accessibility",
      "color": "#0ea5e9"
    }
  ],
  "meta": {
    "total_categories": 7,
    "total_regulations": 15
  }
}
```

### GET /regulations/{id}

**Description**: Get a specific regulation by ID

```json
{
  "data": {
    "id": "cultural-guidelines",
    "category": "cultural",
    "question": "What cultural guidelines should visitors follow?",
    "answer": "Please remove shoes when entering traditional houses. Avoid pointing with your index finger; use an open hand instead. Do not touch sacred objects without permission. During ceremonies, maintain respectful silence and follow guide instructions. Traditional greetings are appreciated but not required.",
    "priority": 4,
    "is_active": true,
    "tags": ["culture", "respect", "traditional", "ceremony"],
    "related_regulations": [
      {
        "id": "photography-rules",
        "question": "What are the photography and videography rules?"
      },
      {
        "id": "general-rules",
        "question": "What are the general rules for visiting Yaro Wora village?"
      }
    ],
    "created_at": "2024-01-10T08:00:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

---

## 🏢 Facilities Page

### GET /facilities

**Description**: Get available facilities and experiences

**Query Parameters**:

- `category`: `accommodation|workshop|culinary|entertainment|activity|wellness|educational|adventure`
- `limit`: number of results per page
- `offset`: pagination offset
- `lang`: `en|id` (language for localized content)
- `sort`: `price|duration|popularity` (sorting options)

```json
{
  "data": [
    {
      "id": "traditional-homestay",
      "name": "Traditional Homestay",
      "name_id": "Homestay Tradisional",
      "description": "Experience authentic village life by staying with local families in traditional Sumba houses. Enjoy home-cooked meals, learn daily customs, and participate in family activities.",
      "description_id": "Rasakan kehidupan desa autentik dengan menginap bersama keluarga lokal di rumah tradisional Sumba. Nikmati makanan rumahan, pelajari adat sehari-hari, dan berpartisipasi dalam kegiatan keluarga.",
      "image_url": "https://api.yarowora.com/images/facilities/homestay.jpg",
      "thumbnail_url": "https://api.yarowora.com/images/facilities/thumbs/homestay.jpg",
      "category": "accommodation",
      "category_id": "akomodasi",
      "highlights": [
        "Traditional Architecture",
        "Family Integration",
        "Home-cooked Meals",
        "Cultural Immersion"
      ],
      "highlights_id": [
        "Arsitektur Tradisional",
        "Integrasi Keluarga",
        "Makanan Rumahan",
        "Pendalaman Budaya"
      ],
      "duration": "1-7 nights",
      "capacity": "2-4 guests per family",
      "price": 350000,
      "currency": "IDR",
      "booking_required": true,
      "advance_booking_days": 3,
      "availability": {
        "days": [
          "monday",
          "tuesday",
          "wednesday",
          "thursday",
          "friday",
          "saturday",
          "sunday"
        ],
        "seasonal_notes": "Available year-round, best during dry season"
      },
      "includes": [
        "Traditional accommodation",
        "3 meals per day",
        "Family activities",
        "Cultural guidance"
      ],
      "includes_id": [
        "Akomodasi tradisional",
        "3 kali makan per hari",
        "Aktivitas keluarga",
        "Panduan budaya"
      ],
      "requirements": ["Respect for local customs", "Basic Indonesian helpful"],
      "requirements_id": [
        "Menghormati adat lokal",
        "Bahasa Indonesia dasar membantu"
      ]
    },
    {
      "id": "weaving-workshop",
      "name": "Traditional Weaving Workshop",
      "name_id": "Workshop Tenun Tradisional",
      "description": "Learn the ancient art of Sumba ikat weaving from master craftswomen. Create your own traditional textile using time-honored techniques passed down through generations.",
      "description_id": "Pelajari seni kuno tenun ikat Sumba dari pengrajin master. Buat tekstil tradisional Anda sendiri menggunakan teknik kuno yang diturunkan melalui generasi.",
      "image_url": "https://api.yarowora.com/images/facilities/weaving.jpg",
      "thumbnail_url": "https://api.yarowora.com/images/facilities/thumbs/weaving.jpg",
      "category": "workshop",
      "category_id": "workshop",
      "highlights": [
        "Hands-on Experience",
        "Master Craftswomen",
        "Take-home Souvenir",
        "Traditional Techniques"
      ],
      "highlights_id": [
        "Pengalaman Langsung",
        "Pengrajin Master",
        "Oleh-oleh",
        "Teknik Tradisional"
      ],
      "duration": "3-4 hours",
      "capacity": "8 participants",
      "price": 180000,
      "currency": "IDR",
      "booking_required": true,
      "advance_booking_days": 1,
      "availability": {
        "days": ["monday", "wednesday", "friday", "saturday"],
        "times": ["09:00", "14:00"]
      },
      "includes": [
        "All materials",
        "Expert instruction",
        "Traditional snacks",
        "Finished textile"
      ],
      "includes_id": [
        "Semua bahan",
        "Instruksi ahli",
        "Camilan tradisional",
        "Tekstil jadi"
      ],
      "requirements": ["No prior experience needed", "Minimum age 8 years"],
      "requirements_id": ["Tidak perlu pengalaman", "Usia minimum 8 tahun"]
    }
  ],
  "meta": {
    "total": 8,
    "categories": [
      {
        "id": "accommodation",
        "name": "Accommodation",
        "name_id": "Akomodasi",
        "count": 1
      },
      {
        "id": "workshop",
        "name": "Workshop",
        "name_id": "Workshop",
        "count": 1
      },
      {
        "id": "culinary",
        "name": "Culinary",
        "name_id": "Kuliner",
        "count": 1
      },
      {
        "id": "entertainment",
        "name": "Entertainment",
        "name_id": "Hiburan",
        "count": 1
      },
      {
        "id": "activity",
        "name": "Activity",
        "name_id": "Aktivitas",
        "count": 1
      },
      {
        "id": "wellness",
        "name": "Wellness",
        "name_id": "Kesehatan",
        "count": 1
      },
      {
        "id": "educational",
        "name": "Educational",
        "name_id": "Edukasi",
        "count": 1
      },
      {
        "id": "adventure",
        "name": "Adventure",
        "name_id": "Petualangan",
        "count": 1
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 20,
      "total_pages": 1
    }
  }
}
```

### GET /facilities/{id}

**Description**: Get detailed information about a specific facility for the modal

```json
{
  "data": {
    "id": "traditional-homestay",
    "name": "Traditional Homestay",
    "name_id": "Homestay Tradisional",
    "description": "Experience authentic village life by staying with local families...",
    "description_id": "Rasakan kehidupan desa autentik dengan menginap bersama keluarga lokal...",
    "full_content": "# Traditional Homestay Experience\n\nStay with local families and experience...",
    "full_content_id": "# Pengalaman Homestay Tradisional\n\nMenginap dengan keluarga lokal...",
    "images": [
      {
        "url": "https://api.yarowora.com/images/facilities/homestay-1.jpg",
        "caption": "Traditional Sumba house exterior",
        "caption_id": "Eksterior rumah tradisional Sumba",
        "is_primary": true
      },
      {
        "url": "https://api.yarowora.com/images/facilities/homestay-2.jpg",
        "caption": "Family dining area",
        "caption_id": "Area makan keluarga",
        "is_primary": false
      }
    ],
    "category": "accommodation",
    "category_id": "akomodasi",
    "highlights": [
      "Traditional Architecture",
      "Family Integration",
      "Home-cooked Meals",
      "Cultural Immersion"
    ],
    "highlights_id": [
      "Arsitektur Tradisional",
      "Integrasi Keluarga",
      "Makanan Rumahan",
      "Pendalaman Budaya"
    ],
    "duration": "1-7 nights",
    "capacity": "2-4 guests per family",
    "price": 350000,
    "currency": "IDR",
    "location": {
      "latitude": -9.6234,
      "longitude": 119.3456,
      "address": "Yaro Wora Village, East Sumba",
      "address_id": "Desa Yaro Wora, Sumba Timur"
    },
    "includes": [
      "Traditional accommodation",
      "3 meals per day",
      "Family activities",
      "Cultural guidance"
    ],
    "includes_id": [
      "Akomodasi tradisional",
      "3 kali makan per hari",
      "Aktivitas keluarga",
      "Panduan budaya"
    ],
    "requirements": ["Respect for local customs", "Basic Indonesian helpful"],
    "requirements_id": [
      "Menghormati adat lokal",
      "Bahasa Indonesia dasar membantu"
    ],
    "what_to_bring": [
      "Comfortable clothing",
      "Personal toiletries",
      "Camera",
      "Open mind"
    ],
    "what_to_bring_id": [
      "Pakaian nyaman",
      "Peralatan mandi pribadi",
      "Kamera",
      "Pikiran terbuka"
    ],
    "booking_policy": {
      "advance_booking_days": 3,
      "cancellation_policy": "Free cancellation up to 48 hours before arrival",
      "cancellation_policy_id": "Pembatalan gratis hingga 48 jam sebelum kedatangan",
      "payment_policy": "50% deposit required upon booking",
      "payment_policy_id": "Deposit 50% diperlukan saat pemesanan",
      "group_bookings": "Special rates available for groups of 6+",
      "group_bookings_id": "Tarif khusus tersedia untuk grup 6+ orang"
    },
    "availability": {
      "days": [
        "monday",
        "tuesday",
        "wednesday",
        "thursday",
        "friday",
        "saturday",
        "sunday"
      ],
      "seasonal_notes": "Available year-round, best during dry season (May-October)",
      "seasonal_notes_id": "Tersedia sepanjang tahun, terbaik saat musim kemarau (Mei-Oktober)"
    },
    "contact_info": {
      "phone": "+62 903 009 909",
      "whatsapp": "+62 903 009 909",
      "email": "homestay@yarowora.com"
    }
  }
}
```

### POST /facilities/{id}/book

**Description**: Book a facility/experience

**Request Body**:

```json
{
  "guest_name": "John Doe",
  "email": "john@example.com",
  "phone": "+1234567890",
  "check_in_date": "2024-02-15",
  "check_out_date": "2024-02-17",
  "participants": 2,
  "special_requirements": "Vegetarian meals",
  "language_preference": "en"
}
```

**Response**:

```json
{
  "success": true,
  "booking_id": "BOOK-2024-001",
  "message": "Booking confirmed successfully",
  "total_price": 700000,
  "currency": "IDR",
  "confirmation_email_sent": true
}
```

---

## 📰 News Page

### GET /news

**Description**: Get news articles and updates

**Query Parameters**:

- `limit`: number of articles per page (default: 12)
- `offset`: pagination offset (default: 0)
- `category`: `culture|tourism|events|environment|business` - filter by news category
- `featured`: `true|false` - get only featured/headline articles
- `search`: search query for title, excerpt, or content
- `lang`: `en|id` (language for localized content)

```json
{
  "data": [
    {
      "id": 1,
      "title": "New Cultural Heritage Center Opens in Yaro Wora Village",
      "title_id": "Pusat Warisan Budaya Baru Dibuka di Desa Yaro Wora",
      "excerpt": "A state-of-the-art cultural heritage center has been inaugurated to preserve and showcase the rich traditions of Sumba's Yaro Wora village",
      "excerpt_id": "Pusat warisan budaya canggih telah diresmikan untuk melestarikan dan menampilkan tradisi kaya desa Yaro Wora Sumba",
      "content": "The Yaro Wora Cultural Heritage Center officially opened...",
      "content_id": "Pusat Warisan Budaya Yaro Wora resmi dibuka...",
      "author": {
        "name": "Maria Tandipati",
        "avatar": "https://api.yarowora.com/images/authors/maria.jpg",
        "bio": "Cultural Heritage Specialist"
      },
      "date_published": "2024-12-15T10:00:00Z",
      "category": "culture",
      "featured_image": "https://api.yarowora.com/images/news/heritage-center.jpg",
      "tags": ["heritage", "culture", "museum", "traditional"],
      "read_time": 5,
      "is_headline": true,
      "view_count": 1250,
      "language": "en"
    }
  ],
  "meta": {
    "total": 25,
    "featured_count": 1,
    "categories": [
      {
        "key": "culture",
        "name": "Culture",
        "name_id": "Budaya",
        "count": 8
      },
      {
        "key": "tourism",
        "name": "Tourism",
        "name_id": "Pariwisata",
        "count": 6
      },
      {
        "key": "events",
        "name": "Events",
        "name_id": "Acara",
        "count": 5
      },
      {
        "key": "environment",
        "name": "Environment",
        "name_id": "Lingkungan",
        "count": 3
      },
      {
        "key": "business",
        "name": "Business",
        "name_id": "Bisnis",
        "count": 3
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 12,
      "total_pages": 3,
      "has_next": true,
      "has_previous": false
    }
  }
}
```

### GET /news/categories

**Description**: Get all available news categories with article counts

```json
{
  "data": [
    {
      "key": "culture",
      "name": "Culture",
      "name_id": "Budaya",
      "description": "Cultural events, traditions, and heritage news",
      "description_id": "Berita acara budaya, tradisi, dan warisan",
      "count": 8,
      "color": "#8a0604",
      "icon": "cultural-heritage"
    },
    {
      "key": "tourism",
      "name": "Tourism",
      "name_id": "Pariwisata",
      "description": "Tourism developments and visitor experiences",
      "description_id": "Perkembangan pariwisata dan pengalaman pengunjung",
      "count": 6,
      "color": "#586d12",
      "icon": "tourism"
    },
    {
      "key": "events",
      "name": "Events",
      "name_id": "Acara",
      "description": "Community events and celebrations",
      "description_id": "Acara komunitas dan perayaan",
      "count": 5,
      "color": "#dc2626",
      "icon": "events"
    },
    {
      "key": "environment",
      "name": "Environment",
      "name_id": "Lingkungan",
      "description": "Environmental conservation and sustainability",
      "description_id": "Konservasi lingkungan dan keberlanjutan",
      "count": 3,
      "color": "#16a34a",
      "icon": "environment"
    },
    {
      "key": "business",
      "name": "Business",
      "name_id": "Bisnis",
      "description": "Local business and economic development",
      "description_id": "Bisnis lokal dan pengembangan ekonomi",
      "count": 3,
      "color": "#0ea5e9",
      "icon": "business"
    }
  ],
  "meta": {
    "total_categories": 5,
    "total_articles": 25
  }
}
```

### GET /news/{id}

**Description**: Get full article content with related articles

**Query Parameters**:

- `lang`: `en|id` (language for localized content)

```json
{
  "data": {
    "id": 1,
    "title": "New Cultural Heritage Center Opens in Yaro Wora Village",
    "title_id": "Pusat Warisan Budaya Baru Dibuka di Desa Yaro Wora",
    "excerpt": "A state-of-the-art cultural heritage center has been inaugurated...",
    "excerpt_id": "Pusat warisan budaya canggih telah diresmikan...",
    "content": "The Yaro Wora Cultural Heritage Center officially opened its doors yesterday, marking a significant milestone in the preservation of Sumba's indigenous culture...",
    "content_id": "Pusat Warisan Budaya Yaro Wora resmi membuka pintunya kemarin...",
    "author": {
      "name": "Maria Tandipati",
      "avatar": "https://api.yarowora.com/images/authors/maria.jpg",
      "bio": "Cultural Heritage Specialist and local journalist covering Sumba traditions",
      "email": "maria@yarowora.com",
      "social_links": {
        "twitter": "@mariatandipati",
        "instagram": "@maria_sumba"
      }
    },
    "date_published": "2024-12-15T10:00:00Z",
    "date_updated": "2024-12-15T15:30:00Z",
    "category": "culture",
    "featured_image": "https://api.yarowora.com/images/news/heritage-center.jpg",
    "image_gallery": [
      {
        "url": "https://api.yarowora.com/images/news/gallery/heritage-1.jpg",
        "caption": "Interior view of the cultural heritage center",
        "alt_text": "Museum interior with traditional artifacts"
      },
      {
        "url": "https://api.yarowora.com/images/news/gallery/heritage-2.jpg",
        "caption": "Traditional textiles on display",
        "alt_text": "Sumba ikat textiles exhibition"
      }
    ],
    "tags": ["heritage", "culture", "museum", "traditional"],
    "read_time": 5,
    "is_headline": true,
    "view_count": 1250,
    "related_articles": [
      {
        "id": 2,
        "title": "Sustainable Tourism Initiative Launched",
        "excerpt": "A comprehensive sustainable tourism program...",
        "featured_image": "https://api.yarowora.com/images/news/sustainability.jpg",
        "date_published": "2024-12-10T10:00:00Z",
        "category": "tourism"
      },
      {
        "id": 3,
        "title": "Annual Pasola Festival Draws International Visitors",
        "excerpt": "This year's Pasola festival attracted thousands...",
        "featured_image": "https://api.yarowora.com/images/news/pasola.jpg",
        "date_published": "2024-12-05T10:00:00Z",
        "category": "events"
      }
    ],
    "seo": {
      "meta_title": "Cultural Heritage Center Opens - Yaro Wora News",
      "meta_description": "Discover the new cultural heritage center in Yaro Wora village, preserving Sumba traditions for future generations",
      "keywords": [
        "yaro wora",
        "heritage center",
        "sumba culture",
        "traditional museum"
      ],
      "canonical_url": "https://yarowora.com/news/heritage-center-opens"
    },
    "language": "en"
  }
}
```

### GET /news/featured

**Description**: Get the current featured/headline article

```json
{
  "data": {
    "id": 1,
    "title": "New Cultural Heritage Center Opens in Yaro Wora Village",
    "excerpt": "A state-of-the-art cultural heritage center...",
    "featured_image": "https://api.yarowora.com/images/news/heritage-center.jpg",
    "author": {
      "name": "Maria Tandipati",
      "avatar": "https://api.yarowora.com/images/authors/maria.jpg"
    },
    "date_published": "2024-12-15T10:00:00Z",
    "category": "culture",
    "read_time": 5,
    "view_count": 1250
  }
}
```

---

## 📞 Contact & Booking

### POST /contact

**Description**: Submit contact form

**Request Body**:

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "+1234567890",
  "subject": "Booking Inquiry",
  "message": "I would like to book a visit for 4 people...",
  "preferred_date": "2024-02-15",
  "visitor_type": "foreigner",
  "visitor_count": {
    "adults": 4,
    "infants": 2
  }
}
```

**Response**:

```json
{
  "success": true,
  "message": "Your inquiry has been submitted successfully. We will contact you within 24 hours.",
  "reference_id": "INQ-2024-001",
  "estimated_response_time": "24 hours"
}
```

### GET /contact-info

**Description**: Get contact information and location details

```json
{
  "data": {
    "address": {
      "street": "Yaro Wora Village",
      "city": "East Sumba",
      "province": "East Nusa Tenggara",
      "country": "Indonesia",
      "postal_code": "87173"
    },
    "coordinates": {
      "latitude": -9.6234,
      "longitude": 119.3456
    },
    "contact": {
      "phones": ["+62 098 940 974", "+62 903 009 909"],
      "email": ["info@yarowora.com", "visit@yarowora.com"],
      "whatsapp": "+62 903 009 909"
    },
    "social_media": {
      "instagram": "@yarowora_official",
      "facebook": "Yaro Wora Tourism",
      "youtube": "Yaro Wora Channel"
    },
    "operating_hours": {
      "monday": "08:00-17:00",
      "tuesday": "08:00-17:00",
      "wednesday": "08:00-17:00",
      "thursday": "08:00-17:00",
      "friday": "08:00-17:00",
      "saturday": "08:00-16:00",
      "sunday": "closed"
    }
  }
}
```

---

## 🌐 Multi-language Support

### GET /content/{lang}

**Description**: Get localized content for the specified language

**Path Parameters**:

- `lang`: `en|id` (English or Indonesian)

All endpoints support language query parameter:

```
GET /attractions?lang=id
GET /news?lang=en
```

**Response includes localized content**:

```json
{
  "data": {
    "title": "Rumah Tradisional Sumba", // Indonesian version
    "description": "Rumah-rumah beratap runcing yang menampilkan arsitektur kuno"
  },
  "language": "id",
  "fallback_language": "en"
}
```

---

## Error Responses

All endpoints return standardized error responses:

```json
{
  "error": true,
  "message": "Resource not found",
  "code": "NOT_FOUND",
  "details": {
    "resource": "attraction",
    "id": "invalid-id"
  },
  "timestamp": "2024-01-15T10:30:00Z"
}
```

**HTTP Status Codes**:

- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 422: Validation Error
- 500: Internal Server Error

---

## Rate Limiting

- Public endpoints: 100 requests per minute per IP
- Upload endpoints: 10 requests per minute per user

## Image Upload Specifications

- Supported formats: JPEG, PNG, WebP
- Maximum file size: 10MB
- Recommended dimensions:
  - Carousel images: 1920x1080px
  - Gallery images: 1200x800px minimum
  - Thumbnails: Auto-generated at 300x200px

## Notes for Backend Implementation

1. **Image Optimization**: Implement automatic image resizing and WebP conversion
2. **Caching**: Implement Redis caching for frequently accessed content
3. **CDN**: Use CDN for image delivery for better performance
4. **Database**: Suggested structure with proper indexing for search functionality
5. **Backup**: Implement automated daily backups for content and images
6. **Monitoring**: Add logging and monitoring for API usage and errors
7. **Security**: Implement proper validation, sanitization, and rate limiting
