# Yaro Wora Tourism Website - Admin API Documentation

This document outlines the admin-only backend endpoints for managing the Yaro Wora tourism website.

## Base URL

```
https://api.yarowora.com/v1/admin
```

## Authentication

- **Required**: Bearer token authentication for all admin endpoints
- **Content-Type**: `application/json`
- **Authorization Header**: `Bearer {admin_token}`

---

## 🏠 Main Page Management

### POST /admin/carousel

**Description**: Create new carousel slide

### PUT /admin/carousel/{id}

**Description**: Update carousel slide

### DELETE /admin/carousel/{id}

**Description**: Delete carousel slide

### POST /admin/attractions

**Description**: Create new attraction

### PUT /admin/attractions/{id}

**Description**: Update attraction

### DELETE /admin/attractions/{id}

**Description**: Delete attraction

### PUT /admin/pricing

**Description**: Update entrance fee pricing

---

## 👤 Profile Page Management

### PUT /admin/profile

**Description**: Update village profile information

**Request Body**:

```json
{
  "title": "Desa Patiala Bawa - Kampung Yaro Wora",
  "description": "Updated village description...",
  "vision": {
    "title": "Visi",
    "content": "Updated vision content..."
  },
  "mission": {
    "title": "Misi",
    "content": "Updated mission content..."
  },
  "objectives": {
    "title": "Tujuan",
    "content": "Updated objectives content..."
  },
  "featured_images": [
    {
      "url": "https://api.yarowora.com/images/profile/village-overview.jpg",
      "caption": "Overview of Yaro Wora village",
      "alt_text": "Traditional houses with mountain backdrop"
    }
  ]
}
```

---

## 🗺️ Destinations Management

### POST /admin/destinations

**Description**: Create new destination

### PUT /admin/destinations/{id}

**Description**: Update destination

### DELETE /admin/destinations/{id}

**Description**: Delete destination

---

## 🖼️ Gallery Management

### POST /admin/gallery

**Description**: Upload new gallery image

### PUT /admin/gallery/{id}

**Description**: Update gallery image information

### DELETE /admin/gallery/{id}

**Description**: Delete gallery image

### POST /admin/gallery/categories

**Description**: Create new gallery category

### PUT /admin/gallery/categories/{id}

**Description**: Update gallery category

### DELETE /admin/gallery/categories/{id}

**Description**: Delete gallery category

---

## 📋 Regulations Management

### POST /admin/regulations

**Description**: Create new regulation

### PUT /admin/regulations/{id}

**Description**: Update existing regulation

### DELETE /admin/regulations/{id}

**Description**: Delete regulation (soft delete)

### POST /admin/regulations/categories

**Description**: Create new regulation category

### PUT /admin/regulations/categories/{id}

**Description**: Update regulation category

### DELETE /admin/regulations/categories/{id}

**Description**: Delete regulation category

---

## 🏢 Facilities Management

### POST /admin/facilities

**Description**: Create new facility/experience

### PUT /admin/facilities/{id}

**Description**: Update facility/experience

### DELETE /admin/facilities/{id}

**Description**: Delete facility/experience

### GET /admin/facilities/{id}/bookings

**Description**: Get bookings for a specific facility

---

## 📰 News Management

### POST /admin/news

**Description**: Create new news article

**Request Body**:

**Description**: Create new news article

### PUT /admin/news/{id}

**Description**: Update news article

### DELETE /admin/news/{id}

**Description**: Delete news article

### POST /admin/news/categories

**Description**: Create new news category

### PUT /admin/news/categories/{id}

**Description**: Update news category

### DELETE /admin/news/categories/{id}

**Description**: Delete news category

---

## 📞 Contact & Booking Management

### GET /admin/contacts

**Description**: Get all contact form submissions

### GET /admin/contacts/{id}

**Description**: Get specific contact submission

### PUT /admin/contacts/{id}

**Description**: Update contact submission status

### GET /admin/bookings

**Description**: Get all facility bookings

### PUT /admin/bookings/{id}

**Description**: Update booking status

---

## 🌐 Content Management

### PUT /admin/contact-info

**Description**: Update contact information and location details

### POST /admin/content/upload

**Description**: Upload images and media files

**Request**: Multipart form data with file upload

**Response**:

```json
{
  "success": true,
  "file_url": "https://api.yarowora.com/images/uploads/filename.jpg",
  "thumbnail_url": "https://api.yarowora.com/images/uploads/thumbs/filename.jpg",
  "file_size": 1024000,
  "dimensions": {
    "width": 1200,
    "height": 800
  }
}
```

---

## 📊 Analytics & Reports

### GET /admin/analytics/visitors

**Description**: Get visitor statistics

### GET /admin/analytics/bookings

**Description**: Get booking analytics

### GET /admin/analytics/content

**Description**: Get content performance metrics

---

## 👥 User Management

### GET /admin/users

**Description**: Get all admin users

### POST /admin/users

**Description**: Create new admin user

### PUT /admin/users/{id}

**Description**: Update admin user

### DELETE /admin/users/{id}

**Description**: Delete admin user

---

## Error Responses

All admin endpoints return standardized error responses:

```json
{
  "error": true,
  "message": "Unauthorized access",
  "code": "UNAUTHORIZED",
  "details": {
    "required_permission": "admin",
    "current_user": "guest"
  },
  "timestamp": "2024-01-15T10:30:00Z"
}
```

**HTTP Status Codes**:

- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden (insufficient permissions)
- 404: Not Found
- 422: Validation Error
- 500: Internal Server Error

---

## Rate Limiting

- Admin endpoints: 1000 requests per minute per authenticated user
- Upload endpoints: 10 requests per minute per user
- Bulk operations: 100 requests per minute per user

---

## Security Notes

1. **Authentication**: All endpoints require valid admin JWT token
2. **Permissions**: Role-based access control (super_admin, content_editor, moderator)
3. **Audit Logging**: All admin actions are logged with user tracking
4. **Input Validation**: Strict validation and sanitization on all inputs
5. **File Upload Security**: Virus scanning and file type validation
6. **Rate Limiting**: Prevents abuse and ensures system stability
