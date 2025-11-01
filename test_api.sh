#!/bin/bash

# Auth Service API Test Script
# This script demonstrates all available endpoints

BASE_URL="http://localhost:8080/api"

echo "================================================"
echo "🧪 Auth Service API Test Suite"
echo "================================================"
echo ""

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Test 1: Health Check
echo -e "${BLUE}📋 Test 1: Health Check${NC}"
echo "GET $BASE_URL/health"
curl -s "$BASE_URL/health" | python3 -m json.tool 2>/dev/null || curl -s "$BASE_URL/health"
echo -e "\n"

# Test 2: Register New User
echo -e "${BLUE}📋 Test 2: Register New User${NC}"
echo "POST $BASE_URL/register"
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/register" \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Test User $(date +%s)\",\"email\":\"test$(date +%s)@example.com\",\"password\":\"testpass123\"}")
echo "$REGISTER_RESPONSE" | python3 -m json.tool 2>/dev/null || echo "$REGISTER_RESPONSE"
TOKEN=$(echo "$REGISTER_RESPONSE" | grep -o '"token":"[^"]*' | cut -d'"' -f4)
EMAIL=$(echo "$REGISTER_RESPONSE" | grep -o '"email":"[^"]*' | cut -d'"' -f4)
echo -e "\n${GREEN}✓ Token saved for authenticated requests${NC}"
echo -e "\n"

# Test 3: Login
echo -e "${BLUE}📋 Test 3: Login${NC}"
echo "POST $BASE_URL/login"
curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\",\"password\":\"testpass123\"}" | python3 -m json.tool 2>/dev/null || \
curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\",\"password\":\"testpass123\"}"
echo -e "\n"

# Test 4: Get Profile (Authenticated)
echo -e "${BLUE}📋 Test 4: Get Profile (Authenticated)${NC}"
echo "GET $BASE_URL/profile"
curl -s -X GET "$BASE_URL/profile" \
  -H "Authorization: Bearer $TOKEN" | python3 -m json.tool 2>/dev/null || \
curl -s -X GET "$BASE_URL/profile" \
  -H "Authorization: Bearer $TOKEN"
echo -e "\n"

# Test 5: Update Profile
echo -e "${BLUE}📋 Test 5: Update Profile${NC}"
echo "PUT $BASE_URL/profile"
curl -s -X PUT "$BASE_URL/profile" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Name","avatar":"https://example.com/avatar.jpg"}' | python3 -m json.tool 2>/dev/null || \
curl -s -X PUT "$BASE_URL/profile" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Name","avatar":"https://example.com/avatar.jpg"}'
echo -e "\n"

# Test 6: Get Updated Profile
echo -e "${BLUE}📋 Test 6: Verify Profile Update${NC}"
echo "GET $BASE_URL/profile"
curl -s -X GET "$BASE_URL/profile" \
  -H "Authorization: Bearer $TOKEN" | python3 -m json.tool 2>/dev/null || \
curl -s -X GET "$BASE_URL/profile" \
  -H "Authorization: Bearer $TOKEN"
echo -e "\n"

# Test 7: Invalid Token
echo -e "${BLUE}📋 Test 7: Invalid Token (Should Fail)${NC}"
echo "GET $BASE_URL/profile with invalid token"
curl -s -X GET "$BASE_URL/profile" \
  -H "Authorization: Bearer invalid-token-12345" | python3 -m json.tool 2>/dev/null || \
curl -s -X GET "$BASE_URL/profile" \
  -H "Authorization: Bearer invalid-token-12345"
echo -e "\n"

# Test 8: Wrong Password
echo -e "${BLUE}📋 Test 8: Login with Wrong Password (Should Fail)${NC}"
echo "POST $BASE_URL/login"
curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\",\"password\":\"wrongpassword\"}" | python3 -m json.tool 2>/dev/null || \
curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\",\"password\":\"wrongpassword\"}"
echo -e "\n"

echo "================================================"
echo -e "${GREEN}✅ All tests completed!${NC}"
echo "================================================"
echo ""
echo "To test account deletion (destructive):"
echo "  curl -X DELETE $BASE_URL/profile -H \"Authorization: Bearer \$TOKEN\""
echo ""
