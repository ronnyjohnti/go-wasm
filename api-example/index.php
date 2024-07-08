<?php

header('Content-Type', 'application/json');

echo json_encode([
  'content' => 'Alguma coisa de conteúdo',
  'message' => 'Aqui temos algo',
  'code' => 2
]);
