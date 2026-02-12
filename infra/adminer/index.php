<?php
use Adminer\Adminer;

if (basename($_SERVER['DOCUMENT_URI'] ?? $_SERVER['REQUEST_URI']) === 'adminer.css' && is_readable('adminer.css')) {
    header('Content-Type: text/css');
    readfile('adminer.css');
    exit;
}

// Enable login without password for SQLite
function adminer_object()
{
    class AdminerSoftware extends Adminer
    {
        function login($login, $password)
        {
            // Always allow login for SQLite (no password needed)
            return true;
        }

        function loginForm()
        {
            // Auto-fill the login form (but don't auto-submit)
            ?>
                        <table cellspacing="0" class="layout">
                            <tr>
                                <th>
                                    System
                                <td>
                                    <input type="hidden" name="auth[driver]" value="sqlite">
                                    SQLite 3
                            <tr>
                                <th>
                                    Database
                                <td>
                                    <input name="auth[db]" value="/data/local.db" autocapitalize="off">
                        </table>
                        <p><input type="submit" value="Login">
                            <?php
        }
    }

    return new AdminerSoftware;
}

include __DIR__ . '/adminer.php';
