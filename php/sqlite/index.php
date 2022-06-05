<?php
// This code is based on example of Learning PHP (O'reilly)
$db = new PDO('sqlite:dinner.db');
$meals = array('breakfast', 'lunch', 'dinner');
if (in_array($_GET['meal'], $meals)) {
  $stmt = $db->prepare('SELECT name, price FROM meals WHERE meal = ?');
  $stmt->execute(array($_GET['meal']));
  $rows = $stmt->fetchAll();
  if (count($rows) === 0) {
    print "No rows found";
    return;
  }
  print '<table><tr><th>Name</th><th>Price</th></tr>';
  foreach ($rows as $row) {
    print "<tr><td>$row[0]</td><td>$row[1]</td></tr>";
  }
  print '</table>';
  return;
}
print "Unknown meal.";
?>
