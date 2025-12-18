LinkedIn Automation Proof-of-Concept (Go + Rod)
Overview

This project is a technical proof-of-concept demonstrating how a LinkedIn-class automation system can be architected using Go and the Rod browser automation library, with a strong emphasis on:

Human-like interaction patterns

Anti-bot detection awareness

Ethical guardrails and defensive automation

Clean, modular Go architecture

⚠️ Important: This project is strictly for educational and evaluation purposes.
It is not intended for production use or for automating LinkedIn or any third-party platform in violation of their Terms of Service.

Key Design Philosophy

The goal of this project is not to bypass platform protections, but to demonstrate:

How such systems are designed responsibly

How automation detects and aborts when security checkpoints (CAPTCHA, challenges) are encountered

How advanced automation logic can be safely demonstrated using a controlled environment

Core Features Implemented
Authentication System

Environment-based credential loading (structure in place)

Login page state detection

CAPTCHA and checkpoint detection with immediate abort

Session cookie persistence and reuse

Search & Targeting (Demo-Safe)

Profile search by job title, company, location

Pagination-ready architecture

Profile URL extraction

Duplicate profile detection

Connection Requests (Demo-Safe)

Profile navigation simulation

Personalized connection notes

Daily rate limiting

Persistent tracking of sent requests

Messaging System (Demo-Safe)

Accepted connection detection (simulated)

Template-based follow-up messages

Dynamic variables (e.g., {{name}}, {{company}})

Message history persistence

Anti-Bot & Stealth Techniques

This project implements 8+ stealth and realism techniques, including all mandatory requirements:

Mandatory

Human-like mouse movement using curved paths

Randomized timing between actions

Browser fingerprint masking (viewport, automation flags)

Additional

Realistic typing simulation

Random scrolling behavior

Mouse hovering and micro-movements

Rate limiting and cooldown enforcement

Defensive CAPTCHA detection and abort logic

Session reuse via cookies

The system is intentionally designed to stop automation immediately when CAPTCHA or security challenges are detected.

Demo Mode (Safe Execution)

To visually demonstrate all features without violating any platform rules, the system supports a Demo Mode using a locally hosted website.

Demo Website

Local HTML pages simulate:

Search results

Profile pages

Connect buttons

Messaging forms

Why Demo Mode?

Avoids Terms of Service violations

Allows full visual demonstration (mouse, typing, clicks)

Clearly separates engineering capability from ethical execution

Project Structure
cmd/app/ → Application entry point
internal/
├── auth/ → Login & state detection
├── browser/ → Stealth browser launcher
├── search/ → Search & deduplication logic
├── connect/ → Connection request handling
├── messaging/ → Follow-up messaging system
├── stealth/ → Mouse, typing, timing realism
├── limiter/ → Rate limiting & throttling
├── storage/ → Cookies, connections, messages
├── models/ → Shared domain models
└── logger/ → Structured logging
demo-site/ → Local demo website

How to Run
Prerequisites

Go 1.20+

Python (for demo site)

Chromium (auto-managed by Rod)

Run Demo Site
cd demo-site
python -m http.server 8080

Run Automation
cd ..
go run ./cmd/app
