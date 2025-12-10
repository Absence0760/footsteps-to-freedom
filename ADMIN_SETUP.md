# Admin Panel Setup

This document explains how to set up and use the admin panel for managing contact form submissions.

## Security Features

The admin panel is now protected with JWT-based authentication:

1. **Login Required**: Users must authenticate before accessing the admin dashboard
2. **Protected API Endpoints**: All admin endpoints require a valid JWT token
3. **No Public Links**: The admin page is not linked from the main navigation
4. **Secure Tokens**: JWT tokens are stored in HTTP-only cookies
5. **Session Management**: Tokens expire after 24 hours

## Configuration

### 1. Set Admin Credentials

Edit your `.env` file and set secure credentials:

```bash
# IMPORTANT: Change these default credentials!
ADMIN_USERNAME=your_admin_username
ADMIN_PASSWORD=your_secure_password

# Also set a strong JWT secret
JWT_SECRET=your_random_jwt_secret_string_here
```

**Note**: The default credentials are `admin/admin123` - **CHANGE THESE IN PRODUCTION!**

### 2. Generate a Secure JWT Secret

For production, generate a strong random secret:

```bash
# On Linux/Mac:
openssl rand -base64 32

# Or use any random string generator
```

## Accessing the Admin Panel

1. Navigate to: `http://your-domain.com/admin.html`
2. Enter your admin credentials
3. You'll be logged in for 24 hours

## Features

### Contact Management

- **View All Submissions**: See all contact form submissions in one place
- **Filter by Status**: Filter contacts by New, Read, or Responded
- **Status Management**: Mark contacts as Read or Responded
- **Delete Submissions**: Remove spam or unwanted submissions
- **Email Links**: Click email addresses to open in your mail client

### API Endpoints

All admin endpoints require authentication:

- `POST /api/admin/login` - Authenticate and get JWT token
- `POST /api/admin/logout` - Clear authentication token
- `GET /api/admin/check` - Verify authentication status
- `GET /api/admin/contacts` - Get all contact submissions
- `GET /api/admin/contacts/:id` - Get specific contact
- `PUT /api/admin/contacts/:id` - Update contact status
- `DELETE /api/admin/contacts/:id` - Delete contact

## Security Best Practices

1. **Change Default Credentials**: Never use `admin/admin123` in production
2. **Use HTTPS**: Always use HTTPS in production (set `ENV=production`)
3. **Strong Passwords**: Use long, random passwords
4. **Rotate JWT Secret**: Change your JWT secret periodically
5. **Monitor Access**: Check your logs for unauthorized access attempts
6. **Environment Variables**: Never commit `.env` files to version control

## Troubleshooting

### Can't Log In

- Check that `ADMIN_USERNAME` and `ADMIN_PASSWORD` are set in `.env`
- Verify the backend server is running
- Check browser console for errors

### Token Expired

- Tokens expire after 24 hours
- Simply log in again to get a new token

### Unauthorized Errors

- Clear browser cookies and log in again
- Check that `JWT_SECRET` matches between sessions
- Verify your token hasn't expired

## Development vs Production

### Development
- Uses default credentials if not set
- Cookies use `Secure: false` (works over HTTP)
- More verbose error messages

### Production
- **Must** set custom credentials
- Cookies use `Secure: true` (requires HTTPS)
- Generic error messages

To enable production mode:
```bash
ENV=production
```
